package main

import (
	"context"
	"errors"
	"time"

	"fast-client/service"
)

// ErrBlocked reports if service is blocked.
var ErrBlocked = errors.New("blocked")

// Service defines external service that can process batches of queue.
type Service interface {
	GetLimits() (n uint64, p time.Duration)
	Process(ctx context.Context, batch service.Batch) error
}

type ContextWithItems struct {
	ctx   context.Context
	errCh chan error
	items []service.Item
}

// Client клиент к сервису
type Client struct {
	service Service

	batchLen uint64
	interval time.Duration
	queue    chan ContextWithItems // Очередь на отправку
}

// NewClient создает клиент к сервису
func NewClient(service Service) *Client {
	c := &Client{service: service, queue: make(chan ContextWithItems)}
	c.batchLen, c.interval = service.GetLimits()

	go c.sendFromQueue()

	return c
}

// sendFromQueue заполняет буфер для эффективной отправки в сервис и отправляет
func (c *Client) sendFromQueue() {
	itemsForSend := make([]service.Item, 0, c.batchLen) // буфер для отправки
	errChans := make([]chan error, 0, c.batchLen)       // errCh элементов, попавших в буфер

	for {
		select {
		case v := <-c.queue:
			if len(itemsForSend)+len(v.items) > int(c.batchLen) { // если не помещаемся в лимит, отправляем то, что есть
				err := c.service.Process(v.ctx, itemsForSend)

				for _, errChan := range errChans { // оповещаем замороженные гоурутины о завершении обработки
					errChan <- err
				}

				time.Sleep(c.interval) // ждем готовности сервиса

				// очищаем буфер и сразу добавляем следующий элемент из очереди
				itemsForSend = v.items
				errChans = []chan error{v.errCh}
			} else {
				itemsForSend = append(itemsForSend, v.items...)
				errChans = append(errChans, v.errCh)
			}

		default: // если в очереди больше нет элементов для отправки, отправляем
			err := c.service.Process(context.Background(), itemsForSend)

			for _, errChan := range errChans {
				errChan <- err
			}

			// Очищаем буфер. Очередь пуста
			itemsForSend = nil
			errChans = nil
		}
	}

}

// Send ставит данные в очередь и отправляет в сервис
func (c *Client) Send(ctx context.Context, items []service.Item) error {
	errCh := make(chan error)
	c.queue <- ContextWithItems{ctx: ctx, items: items, errCh: errCh} // отправляем элементы в очередь

	select {
	case err := <-errCh: // ждем, когда наши данные отправят из очереди в сервис
		return err
	case <-ctx.Done():
		return nil
	}
}

// Close закрывает клиент
func (c *Client) Close() {
	close(c.queue)
}

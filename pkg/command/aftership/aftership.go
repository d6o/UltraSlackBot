package afterShip

import (
	"fmt"
	"strings"

	"github.com/jessfraz/ship/aftership"
)

const (
	createFormat = "Created shipment for tracking number: %s"
	getFormat    = "%s (%s) - %s"
)

type (
	afterShip struct {
		client *aftership.Client
	}
)

func newAfterShip(key string) *afterShip {
	return &afterShip{
		client: aftership.New(key),
	}
}

func (as *afterShip) create(tracking string) (string, error) {
	_, err := as.client.PostTracking(
		aftership.Tracking{
			TrackingNumber: tracking,
		},
	)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(createFormat, tracking), nil
}

func (as *afterShip) get(tracking string) (string, error) {
	slug, err := as.slugFromTracking(tracking)
	if err != nil {
		return "", err
	}

	t, err := as.client.GetTracking(
		aftership.Tracking{
			TrackingNumber: tracking,
			Slug:           slug,
		},
	)
	if err != nil {
		return "", err
	}

	return as.formatTracking(t), nil
}

func (as *afterShip) slugFromTracking(tracking string) (string, error) {
	courier, err := as.client.DetectCourier(
		aftership.Tracking{
			TrackingNumber: tracking,
		},
	)
	return courier.Slug, err
}

func (as *afterShip) formatTracking(t aftership.Tracking) string {
	return fmt.Sprintf(getFormat, t.TrackingNumber, strings.ToUpper(t.Slug), t.Tag)
}

func (as *afterShip) list() (string, error) {
	trackingList, err := as.client.GetTrackings()
	if err != nil {
		return "", err
	}

	var msgList []string

	for _, t := range trackingList {
		msgList = append(msgList, as.formatTracking(t))
	}

	return strings.Join(msgList, " "), nil
}

func (as *afterShip) remove(tracking string) (string, error) {
	slug, err := as.slugFromTracking(tracking)
	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf("Deleted shipment for tracking number: %s", tracking)
	return msg, as.client.DeleteTracking(
		aftership.Tracking{
			Slug:           slug,
			TrackingNumber: tracking,
		},
	)
}

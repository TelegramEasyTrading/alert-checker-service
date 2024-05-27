package model

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestUnmarshalAlert(t *testing.T) {
	expectedAlert := &Alert{
		Id:        "1",
		Symbol:    "AAPL",
		Value:     100,
		Condition: Condition_PRICE_ABOVE,
	}

	marshalled, err := proto.Marshal(expectedAlert)
	require.NoError(t, err)
	alert := &Alert{}
	err = proto.Unmarshal(marshalled, alert)
	require.NoError(t, err)
	require.Equal(t, expectedAlert.Value, alert.Value)

}

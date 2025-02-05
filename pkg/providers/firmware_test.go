package providers

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestBuildBoard(t *testing.T) {
	firmwareProvider := NewFirmwareProvider()
	config := FirmwareConfig{
		BoardType:                      "BOARD_YARDFORCE500",
		PanelType:                      "PANEL_TYPE_YARDFORCE_500_CLASSIC",
		DebugType:                      "NONE",
		MaxChargeCurrent:               28,
		LimitVoltage150MA:              29,
		MaxChargeVoltage:               29,
		BatChargeCutoffVoltage:         29,
		OneWheelLiftEmergencyMillis:    10000,
		BothWheelsLiftEmergencyMillis:  100,
		TiltEmergencyMillis:            1000,
		StopButtonEmergencyMillis:      100,
		PlayButtonClearEmergencyMillis: 1000,
		ExternalImuAcceleration:        true,
		ExternalImuAngular:             true,
		MasterJ18:                      true,
		MaxMps:                         0.6,
	}
	res, err := firmwareProvider.buildBoardHeader("./asserts/board.h.template", config)
	assert.NoError(t, err)
	file, err := os.ReadFile("./asserts/board.h")
	assert.Equal(t, string(file), string(res))
}

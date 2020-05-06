package main

import (
	"errors"
	"io/ioutil"
	"strconv"
	"time"
)

// Constants describing car's plug state
const (
	PlugStateNotPlugged     = 0
	PlugStatePartialPlugged = 1
	PlugStatePlugged        = 2
)

// Constants describing car's charge mode
const (
	ChargeModeNotCharging = 0
	ChargeModeL1          = 1
	ChargeModeL2          = 2
	ChargeModeL3          = 3
)

// Constants describing car's current gear
const (
	GearNotReady = 0
	GearPark     = 1
	GearReverse  = 2
	GearNeutral  = 3
	GearDrive    = 4
	GearBEco     = 7
)

// DataLine represents a single line in LeafSpy's CSV logfile
type DataLine struct {
	DateTime    time.Time
	Loc         Location
	Speed       int
	Gids        int
	SOC         int
	AHr         int
	PackVolts   float64
	PackAmps    float64
	MaxCP       int
	MinCP       int
	AvgCP       int
	DiffCP      int
	Judgment    int
	PackT1F     float64
	PackT1C     float64
	PackT2F     float64
	PackT2C     float64
	PackT3F     float64
	PackT3C     float64
	PackT4F     float64
	PackT4C     float64
	CP          map[int]int
	Bat12VAmp   float64
	VIN         string
	Hx          float64
	Bat12VVolt  float64
	OdoKm       float64
	QC          int
	L1L2        int
	TPFL        float64
	TPFR        float64
	TPRR        float64
	TPRL        float64
	AmbientF    float64
	SOH         float64
	RegenWh     int
	BLevel      int
	EpochTime   float64
	MotorPwr    int
	AuxPwr      int
	ACPwr       int
	ACComp      int
	EstPwrAC    int
	EstPwrHtr   int
	PlugState   int
	ChargeMode  int
	OBCOutPwr   int
	Gear        int
	HVolt1      float64
	HVolt2      float64
	GPSStatus   string
	PowerSW     bool
	BMS         bool
	OBC         bool
	Debug       string
	MotorTemp   int
	Inverter2   int
	Inverter4   int
	Speed1      float64
	Speed2      float64
	WiperStatus string
	TorqueNm    float64
}

// Marshal something
func (line *DataLine) Marshal(fields []string) (err error) {
	// Check if we have required number of fields
	if len(fields) != 159 {
		return errors.New("Number of columns in CSV file not appropriate")
	}

	// Date/Time
	// Mon Jan 2 15:04:05 -0700 MST 2006
	line.DateTime, err = time.Parse("2006/01/02 15:04:05", fields[0])
	if err != nil {
		return err
	}

	// Loc
	line.Loc, err = parseLocation(fields[1], fields[2], fields[3])
	if err != nil {
		return err
	}

	// Speed
	line.Speed, err = strconv.Atoi(fields[4])
	if err != nil {
		return err
	}

	// Gids
	line.Gids, err = strconv.Atoi(fields[5])
	if err != nil {
		return err
	}

	// SOC
	line.SOC, err = strconv.Atoi(fields[6])
	if err != nil {
		return err
	}

	// AHr
	line.AHr, err = strconv.Atoi(fields[7])
	if err != nil {
		return err
	}

	// PackVolts
	line.PackVolts, err = strconv.ParseFloat(fields[8], 64)
	if err != nil {
		return err
	}

	// PackAmps
	line.PackAmps, err = strconv.ParseFloat(fields[9], 64)
	if err != nil {
		return err
	}

	// MaxCP
	line.MaxCP, err = strconv.Atoi(fields[10])
	if err != nil {
		return err
	}

	// MinCP
	line.MinCP, err = strconv.Atoi(fields[11])
	if err != nil {
		return err
	}

	// AvgCP
	line.AvgCP, err = strconv.Atoi(fields[12])
	if err != nil {
		return err
	}

	// DiffCP
	line.DiffCP, err = strconv.Atoi(fields[13])
	if err != nil {
		return err
	}

	// Judgment
	line.Judgment, err = strconv.Atoi(fields[14])
	if err != nil {
		return err
	}

	// PackT1F
	if fields[15] != "none" {
		line.PackT1F, err = strconv.ParseFloat(fields[15], 64)
		if err != nil {
			return err
		}
	}

	// PackT1C
	if fields[16] != "none" {
		line.PackT1C, err = strconv.ParseFloat(fields[16], 64)
		if err != nil {
			return err
		}
	}

	// PackT2F
	if fields[17] != "none" {
		line.PackT2F, err = strconv.ParseFloat(fields[17], 64)
		if err != nil {
			return err
		}
	}

	// PackT2C
	if fields[18] != "none" {
		line.PackT2C, err = strconv.ParseFloat(fields[18], 64)
		if err != nil {
			return err
		}
	}

	// PackT3F
	if fields[19] != "none" {
		line.PackT3F, err = strconv.ParseFloat(fields[19], 64)
		if err != nil {
			return err
		}
	}

	// PackT3C
	if fields[20] != "none" {
		line.PackT3C, err = strconv.ParseFloat(fields[20], 64)
		if err != nil {
			return err
		}
	}

	// PackT4F
	if fields[21] != "none" {
		line.PackT4F, err = strconv.ParseFloat(fields[21], 64)
		if err != nil {
			return err
		}
	}

	// PackT4C
	if fields[22] != "none" {
		line.PackT4C, err = strconv.ParseFloat(fields[22], 64)
		if err != nil {
			return err
		}
	}

	// CPxx
	line.CP = make(map[int]int)
	for i := 1; i <= 96; i++ {
		line.CP[i], err = strconv.Atoi(fields[22+i]) // 22 is previous line
	}

	// Bat12VAmp
	line.Bat12VAmp, err = strconv.ParseFloat(fields[119], 64)
	if err != nil {
		return err
	}

	// VIN
	line.VIN = fields[120]

	// Hx
	line.Hx, err = strconv.ParseFloat(fields[121], 64)
	if err != nil {
		return err
	}

	// Bat12VVolt
	line.Bat12VVolt, err = strconv.ParseFloat(fields[122], 64)
	if err != nil {
		return err
	}

	// OdoKm
	line.OdoKm, err = strconv.ParseFloat(fields[123], 64)
	if err != nil {
		return err
	}

	// QC
	line.QC, err = strconv.Atoi(fields[124])
	if err != nil {
		return err
	}

	// L1L2
	line.L1L2, err = strconv.Atoi(fields[125])
	if err != nil {
		return err
	}

	// TPFL
	line.TPFL, err = strconv.ParseFloat(fields[126], 64)
	if err != nil {
		return err
	}

	// TPFR
	line.TPFR, err = strconv.ParseFloat(fields[127], 64)
	if err != nil {
		return err
	}

	// TPRR
	line.TPRR, err = strconv.ParseFloat(fields[128], 64)
	if err != nil {
		return err
	}

	// TPRL
	line.TPRL, err = strconv.ParseFloat(fields[129], 64)
	if err != nil {
		return err
	}

	// AmbientF
	line.AmbientF, err = strconv.ParseFloat(fields[130], 64)
	if err != nil {
		return err
	}

	// SOH
	line.SOH, err = strconv.ParseFloat(fields[131], 64)
	if err != nil {
		return err
	}

	// RegenWh
	line.RegenWh, err = strconv.Atoi(fields[132])
	if err != nil {
		return err
	}

	// BLevel
	line.BLevel, err = strconv.Atoi(fields[133])
	if err != nil {
		return err
	}

	// EpochTime
	line.EpochTime, err = strconv.ParseFloat(fields[134], 64)
	if err != nil {
		return err
	}

	// MotorPwr
	line.MotorPwr, err = strconv.Atoi(fields[135])
	if err != nil {
		return err
	}

	// AuxPwr
	line.AuxPwr, err = strconv.Atoi(fields[136])
	if err != nil {
		return err
	}
	line.AuxPwr *= 100 // Original data in 100w increments

	// ACPwr
	line.ACPwr, err = strconv.Atoi(fields[137])
	if err != nil {
		return err
	}
	line.ACPwr *= 250 // Original data in 250w increments

	// ACComp
	line.ACComp, err = strconv.Atoi(fields[138])
	if err != nil {
		return err
	}

	// EstPwrAC
	line.EstPwrAC, err = strconv.Atoi(fields[139])
	if err != nil {
		return err
	}
	line.EstPwrAC *= 50 // Original data in 50w increments

	// EstPwrHtr
	line.EstPwrHtr, err = strconv.Atoi(fields[140])
	if err != nil {
		return err
	}
	line.EstPwrHtr *= 250 // Original data in 250w increments

	// PlugState
	line.PlugState, err = strconv.Atoi(fields[141])
	if err != nil {
		return err
	}

	// ChargeMode
	line.ChargeMode, err = strconv.Atoi(fields[142])
	if err != nil {
		return err
	}

	// OBCOutPwr
	line.OBCOutPwr, err = strconv.Atoi(fields[143])
	if err != nil {
		return err
	}

	// Gear
	line.Gear, err = strconv.Atoi(fields[144])
	if err != nil {
		return err
	}

	// HVolt1
	line.HVolt1, err = strconv.ParseFloat(fields[145], 64)
	if err != nil {
		return err
	}

	// HVolt2
	line.HVolt2, err = strconv.ParseFloat(fields[146], 64)
	if err != nil {
		return err
	}

	// GPSStatus
	line.GPSStatus = fields[147]

	// PowerSW
	line.PowerSW, err = strconv.ParseBool(fields[148])
	if err != nil {
		return nil
	}

	// BMS
	line.BMS, err = strconv.ParseBool(fields[149])
	if err != nil {
		return nil
	}

	// OBC
	line.OBC, err = strconv.ParseBool(fields[150])
	if err != nil {
		return nil
	}

	// Debug
	line.Debug = fields[151]

	// MotorTemp
	line.MotorTemp, err = strconv.Atoi(fields[152])
	if err != nil {
		return err
	}
	line.MotorTemp -= 40 // Original data C+40

	// Inverter2
	line.Inverter2, err = strconv.Atoi(fields[153])
	if err != nil {
		return err
	}
	line.Inverter2 -= 40 // Original data C+40

	// Inverter4
	line.Inverter4, err = strconv.Atoi(fields[154])
	if err != nil {
		return err
	}
	line.Inverter4 -= 40 // Original data C+40

	// Speed1
	line.Speed1, err = strconv.ParseFloat(fields[155], 64)
	if err != nil {
		return err
	}
	line.Speed1 /= 100 // Original data "2495" for 24.95 km/h

	// Speed2
	line.Speed2, err = strconv.ParseFloat(fields[156], 64)
	if err != nil {
		return err
	}
	line.Speed2 /= 100 // Original data "2495" for 24.95 km/h

	// WiperStatus
	line.WiperStatus = fields[157]

	// TorqueNm
	line.TorqueNm, err = strconv.ParseFloat(fields[158], 64)
	if err != nil {
		return err
	}

	return nil
}

func getFiles(root string) (files []string, err error) {
	fpaths, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}

	for _, file := range fpaths {
		files = append(files, root+file.Name())
	}

	return files, nil

}

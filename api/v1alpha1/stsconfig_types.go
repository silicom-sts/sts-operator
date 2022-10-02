/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type StsGnssSpec struct {
	//
	// Enable/disable GPS
	//
	// Valid range 0-1
	//
	//0 - Disable GPS
	//
	//1 - Enable GPS (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssSigGpsEn int `json:"gnssSigGpsEn,omitempty"`

	//
	// Enable/disable GPS L1C/A
	// Valid range 0-1
	//
	//     0 - Disable GPS L1C/A
	//
	//     1 - Enable GPS L1C/A (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssSigGpsL1CAEn int `json:"gnssSigGpsL1CAEn,omitempty"`

	//
	// Enable/disable GPS L2C
	// Valid range 0-1
	//
	//     0 - Disable GPS L2C
	//
	//     1 - Enable GPS L2C (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssSigGpsL2CEn int `json:"gnssSigGpsL2CEn,omitempty"`

	//
	// Enable/disable SBAS
	// Valid range 0-1
	//
	//     0 - Disable SBAS
	//
	//     1 - Enable SBAS (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssSigSBASEn int `json:"gnssSigSBASEn,omitempty"`

	//
	// Enable/disable SBAS L1C/A
	// Valid range 0-1
	//
	//     0 - Disable SBAS L1C/A (default)
	//
	//     1 - Enable SBAS L1C/A
	//
	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssSigSBASL1CAEn int `json:"gnssSigSBASL1CAEn,omitempty"`

	//
	// Enable/disable Galileo
	// Valid range 0-1
	//
	//     0 - Disable Galileo
	//
	//     1 - Enable Galileo (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssSigGalEn int `json:"gnssSigGalEn,omitempty"`

	//
	// Enable/disable Galileo E1
	// Valid range 0-1
	//
	//     0 - Disable Galileo E1
	//
	//     1 - Enable Galileo E1 (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssSigGalE1En int `json:"gnssSigGalE1En,omitempty"`

	//
	// Enable/disable Galileo E5b
	// Valid range 0-1
	//
	//     0 - Disable Galileo E5b
	//
	//     1 - Enable Galileo E5b (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssSigGalE5BEn int `json:"gnssSigGalE5BEn,omitempty"`

	//
	// Enable/disable BeiDou
	// Valid range 0-1
	//
	//     0 - Disable BeiDou
	//
	//     1 - Enable BeiDou (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssSigBDSEn int `json:"gnssSigBDSEn,omitempty"`

	//
	// Enable/disable BeiDou B1I
	// Valid range 0-1
	//
	//     0 - Disable BeiDou B1I
	//
	//     1 - Enable BeiDou B1I (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssSigBDSB1En int `json:"gnssSigBDSB1En,omitempty"`

	//
	// Enable/disable BeiDou B2I
	// Valid range 0-1
	//
	//     0 - Disable BeiDou B2I
	//
	//     1 - Enable BeiDou B2I (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssSigBDSB2En int `json:"gnssSigBDSB2En,omitempty"`

	//
	// Enable/disable QZSS
	// Valid range 0-1
	//
	//     0 - Disable QZSS
	//
	//     1 - Enable QZSS (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssSigQZSSEn int `json:"gnssSigQZSSEn,omitempty"`

	//
	// Enable/disable QZSS L1C/A
	// Valid range 0-1
	//
	//     0 - Disable QZSS L1C/A
	//
	//     1 - Enable QZSS L1C/A (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssSigQZSSL1CAEn int `json:"gnssSigQZSSL1CAEn,omitempty"`

	//
	// Enable/disable QZSS L1S
	// Valid range 0-1
	//
	//     0 - Disable QZSS L1S (default)
	//
	//     1 - Enable QZSS L1S
	//
	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssSigQZSSL1SEn int `json:"gnssSigQZSSL1SEn,omitempty"`

	//
	// Enable/disable QZSS L2C
	// Valid range 0-1
	//
	//     0 - Disable QZSS L2C
	//
	//     1 - Enable QZSS L2C (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssSigQZSSL2CEn int `json:"gnssSigQZSSL2CEn,omitempty"`

	//
	// Enable/disable GLONASS
	// Valid range 0-1
	//
	//     0 - Disable GLONASS
	//
	//     1 - Enable GLONASS (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssSigGLOEn int `json:"gnssSigGLOEn,omitempty"`

	//
	// Enable/disable GLONASS L1
	// Valid range 0-1
	//
	//     0 - Disable GLONASS L1
	//
	//     1 - Enable GLONASS L1 (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssSigGLOL1En int `json:"gnssSigGLOL1En,omitempty"`

	//
	// Enable/disable GLONASS L2
	// Valid range 0-1
	//
	//     0 - Disable GLONASS L2
	//
	//     1 - Enable GLONASS L2 (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssSigGLOL2En int `json:"gnssSigGLOL2En,omitempty"`

	//
	// ********************************* CFG-TP- Time Pulse Configuration ******************************
	//

	//
	// Antenna cable delay settings (nsec)
	// Valid range +-50,000,000
	//
	//     N - nanoseconds
	//
	// +kubebuilder:default=50
	// +kubebuilder:validation:Minimum=50
	// +kubebuilder:validation:Minimum=-50000000
	// +kubebuilder:validation:Maximum=50000000
	// +kubebuilder:validation:Required
	GnssCableDelay int `json:"gnssCableDelay,omitempty"`

	//
	// Time Pulse is interpreted as Period or Frequency
	// Valid range 0-1
	//
	//     0 - Period
	//
	//     1 - Frequency (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssPulseDef int `json:"gnssPulseDef,omitempty"`

	//
	// Time Pulse Length is interpreted as Pulse Ratio (%) or Length (nsec)
	// Valid range 0-1
	//
	//     0 - Pulse Ratio (default)
	//
	//     1 - Length
	//
	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssPulseLenDef int `json:"gnssPulseLenDef,omitempty"`

	//
	// Enable/disable the first Time Pulse
	// Valid range 0-1
	//
	//     0 - Disable  (default)
	//
	//     1 - Enable
	//
	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssTP1En int `json:"gnssTP1En,omitempty"`

	//
	// Set Time Pulse Frequency (Hz) for the first Time Pulse
	// Valid range
	//
	//     1 - (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssFreqTP1 int `json:"gnssFreqTP1,omitempty"`

	//
	// Set Time pulse Frequency (Hz), when locked to GNSS time, for the first Time Pulse
	// Valid range
	//
	//     1 - (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssFreqLockTP1 int `json:"gnssFreqLockTP1,omitempty"`

	//
	// Use locked parameters, when possible, for the first Time Pulse
	// Valid range 0-1
	//
	//     0 - Disable
	//
	//     1 - Enable (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssUseLockTP1 int `json:"gnssUseLockTP1,omitempty"`

	//
	// Set Time Pulse Duty Cycle (%) for the first Time Pulse
	// Valid range 0-100
	//
	//     0 - (default)
	//
	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssDutyTP1 int `json:"gnssDutyTP1,omitempty"`

	//
	// Set Time Pulse Duty Cycle (%), when locked to GNSS time, for the first Time Pulse
	// Valid range 0-100
	//
	//     10 - (default)
	//
	// +kubebuilder:default:=10
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:validation:Required
	GnssDutyLockTP1 int `json:"gnssDutyLockTP1,omitempty"`

	//
	// Enable/disable the second Time Pulse (10 MHz)
	// Valid range 0-1
	//
	//     0 - Disable  (default)
	//
	//     1 - Enable (default)
	//
	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssTP2En int `json:"gnssTP2En,omitempty"`

	//
	// Set Time Pulse Frequency (Hz) for the second Time Pulse
	// Valid range
	//
	//     10000000 - (default)
	//
	// +kubebuilder:default:=10000000
	// +kubebuilder:validation:Required
	GnssFreqTP2 int `json:"gnssFreqTP2,omitempty"`

	//
	// Set Time pulse Frequency (Hz), when locked to GNSS time, for the second Time Pulse
	// Valid range
	//
	//     10000000 - (default)
	//
	// +kubebuilder:default:=10000000
	// +kubebuilder:validation:Required
	GnssFreqLockTP2 int `json:"gnssFreqLockTP2,omitempty"`

	//
	// Use locked parameters, when possible, for the second Time Pulse
	// Valid range 0-1
	//
	//     0 - Disable
	//
	//     1 - Enable (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssUseLockTP2 int `json:"gnssUseLockTP2,omitempty"`

	//
	// Set Time Pulse Duty Cycle (%) for the second Time Pulse
	// Valid range 0-100
	//
	//     0 - (default)
	//
	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:validation:Required
	GnssDutyTP2 int `json:"gnssDutyTP2,omitempty"`

	//
	// Set Time Pulse Duty Cycle (%), when locked to GNSS time, for the second Time Pulse
	// Valid range 0-100
	//
	//     50 - (default)
	//
	// +kubebuilder:default:=50
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	// +kubebuilder:validation:Required
	GnssDutyLockTP2 int `json:"gnssDutyLockTP2,omitempty"`

	//
	// ********************** CFG-NAVSPG- Standard Precision Navigation Configuration ******************
	//

	//
	// Minimum elevation level (degrees)
	// Valid range
	//
	//     N - degrees
	//
	// +kubebuilder:default:=5
	// +kubebuilder:validation:Required
	GnssMinElev int `json:"gnssMinElev,omitempty"`

	//
	// Minimum satellite signal level for navigation (	Hz)
	// Valid range
	//
	//     N - dBHz
	//
	// +kubebuilder:default:=9
	// +kubebuilder:validation:Required
	GnssMinSatSig int `json:"gnssMinSatSig,omitempty"`

	//
	// ************** CFG-INFMSG- Information Message Configuration **************
	//

	//
	// Information Message Flags for NMEA Protocol on the USB (bitmask)
	// Valid range 0-31
	//
	//     Bit 0 - Enable (default) / Disable Error Information Messages
	//
	//     Bit 1 - Enable (default) / Disable Warning Information Messages
	//
	//     Bit 2 - Enable (default) / Disable Notice Information Messages
	//
	//     Bit 3 - Enable/Disable (default) Test Information Messages
	//
	//     Bit 4 - Enable/Disable (default) Debug Information Messages
	//
	// +kubebuilder:default:=7
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=31
	// +kubebuilder:validation:Required
	GnssMsgNmeaUsb int `json:"gnssMsgNmeaUsb,omitempty"`

	//
	// Information Message Flags for UBX Protocol on the USB (bitmask)
	// Valid range 0-31
	//
	//     Bit 0 - Enable (default) / Disable Error Information Messages
	//
	//     Bit 1 - Enable (default) / Disable Warning Information Messages
	//
	//     Bit 2 - Enable (default) / Disable Notice Information Messages
	//
	//     Bit 3 - Enable/Disable (default) Test Information Messages
	//
	//     Bit 4 - Enable/Disable (default) Debug Information Messages
	//
	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=31
	// +kubebuilder:validation:Required
	GnssMsgUbxUsb int `json:"gnssMsgUbxUsb,omitempty"`

	//
	// *********************** CFG-ITFM- Jamming/Interference Monitor Configuration ********************
	//

	//
	// Enable/disable interference detection
	// Valid range 0-1
	//
	//     0 - Disable interference detection
	//
	//     1 - Enable interference detection (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssIntfDetect int `json:"gnssIntfDetect,omitempty"`

	//
	// Antenna setting
	// Valid range 0-2
	//
	//     0 - Unknown
	//
	//     1 - Passive
	//
	//     2 - Active
	//
	// +kubebuilder:default:=2
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=2
	// +kubebuilder:validation:Required
	GnssAntSet int `json:"gnssAntSet,omitempty"`

	//
	// GNSS CW Jamming Detection Threshold (dB)
	// Valid range 0-255
	//
	//
	// +kubebuilder:default:=50
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:validation:Required
	GnssCwTh int `json:"gnssCwTh,omitempty"`

	//
	// **************************** CFG-TMODE- Survey-in Time Mode Configuration ***********************
	//

	//
	// Receiver Time Mode
	// Valid range 0-2
	//
	//     0 - Disabled
	//
	//     1 - Survey-in (default)
	//
	//     2 - Fixed Mode (true ARP position information required)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=2
	// +kubebuilder:validation:Required
	GnssRecvTMode int `json:"gnssRecvTMode,omitempty"`

	//
	// Survey-in Minimum Duration (sec)
	// Valid range
	//
	//     120 - (default)
	//
	// +kubebuilder:default:=120
	// +kubebuilder:validation:Required
	GnssSvinMinDur int `json:"gnssSvinMinDur,omitempty"`

	//
	// Survey-in Position Accuracy Limit (mm)
	// Valid range
	//
	//     100000 - (default)
	//
	// +kubebuilder:default:=100000
	// +kubebuilder:validation:Required
	GnssSvinAccLimit int `json:"gnssSvinAccLimit,omitempty"`

	//
	// **************************** GNSS Clock Out Configuration ***********************
	//

	//
	// GNSS Lock Mode
	// Valid range 0-1
	//
	//     0 - Manual
	//
	//     1 - Auto (default)
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssLockMode int `json:"gnssLockMode,omitempty"`

	//
	// GNSS Lock Threshold (nsec)
	// Valid range 0-10000
	//
	//     100 - (default)
	//
	// +kubebuilder:default:=100
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=10000
	// +kubebuilder:validation:Required
	GnssLockTh int `json:"gnssLockTh,omitempty"`

	//
	// Enable/disable the Clock Out
	// Valid range 0-1
	//
	//     0 - Disable  (default)
	//
	//     1 - Enable
	//
	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	GnssClockOutEn int `json:"gnssClockOutEn,omitempty"`
}

// StsConfigSpec defines the desired state of StsConfig
type StsConfigSpec struct {

	// +kubebuilder:validation:Required
	Interfaces   []StsInterfaceSpec `json:"interfaces"`
	NodeSelector map[string]string  `json:"nodeSelector,omitempty"`

	// +kubebuilder:default={"gnssSigGpsEn":1}
	// +kubebuilder:validation:Required
	GnssSpec StsGnssSpec `json:"gnssSpec,omitempty"`

	// +kubebuilder:validation:Enum=T-GM.8275.1;T-BC-8275.1;T-TSC.8275.1;T-GM.8275.2;T-BC-P-8275.2;
	// +kubebuilder:default:="T-GM.8275.1"
	// +kubebuilder:validation:Required
	//Telecom G8275 Profile
	//
	// T-BC-8275.1
	//
	// T-GM.8275.1  (default)
	//
	// T-TSC.8275.1
	//
	// T-GM.8275.2
	//
	// T-BC-P-8275.2
	Mode string `json:"mode,omitempty"`

	// +kubebuilder:default:="openshift-operators"
	// +kubebuilder:validation:Required
	Namespace string `json:"namespace,omitempty"`

	// +kubebuilder:default:=2
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2
	//Set 1PPS Connector Mode
	//
	//1 - PPS IN
	//
	//2 - PPS OUT (default)
	//
	// +kubebuilder:validation:Required
	ModePPS int `json:"modePPS,omitempty"`

	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3
	// Set PPS OUT Source
	//
	// 1 - PLL (default)
	//
	// 2 - GPS
	//
	// 3 - IN
	// +kubebuilder:validation:Required
	SrcPPS int `json:"srcPPS,omitempty"`

	// +kubebuilder:default:=2
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3
	// Set 10MHz Connector Mode
	//
	// 1 - 10MHz IN
	//
	// 2 - 10MHz OUT (default)
	//
	// 3 - PPS OUT
	// +kubebuilder:validation:Required
	Mode10MHz int `json:"mode10MHz,omitempty"`

	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3
	// Set 10MHz OUT Source
	//
	// 1 - PLL (default)
	//
	// 2 - GPS
	//
	// 3 - IN
	// +kubebuilder:validation:Required
	Src10MHz int `json:"src10MHz,omitempty"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=7
	// Set SyncE Recovery Clock Port
	// +kubebuilder:validation:Required
	SynceRecClkPort int `json:"synceRecClkPort,omitempty"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	//Enable/disable Phy Leds Control Switch
	//
	//0 - disable Phy Leds Control Switch (default)
	//
	//1 - enable Phy Leds Control Switch
	// +kubebuilder:validation:Required
	PhyLedsCtl int `json:"phyLedsCtl,omitempty"`

	//
	// Event type mask
	// Valid range 0-3 (3 is default)
	//
	//Bit 0 - Alarm
	//
	//Bit 1 - Info
	//
	// +kubebuilder:default:=3
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=3
	// +kubebuilder:validation:Required
	EventMask int `json:"eventMask,omitempty"`

	//
	// Enable/disable write tAcc to log
	// Valid range 0-1
	//
	//0 - Disable write log (default)
	//
	//1 - Enable write log
	//
	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	TaccLogEn int `json:"taccLogEn,omitempty"`

	//
	// Set Holdover Quality Categories
	// Valid range 1-3
	//
	//1 - FREQ_CAT1 (default)
	//
	//2 - FREQ_CAT2
	//
	//3 - FREQ_CAT3
	//
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Required
	FreqCat int `json:"freqCat,omitempty"`

	//
	// Set time within Holdover Specification
	// Valid range 0-43200 secs (12 hours)
	//
	//     14400 - (default)
	//
	// +kubebuilder:default:=14400
	// +kubebuilder:validation:Required
	HoSpecDuration int `json:"hoSpecDuration,omitempty"`

	//
	// Set SyncE Recovery Clock Mode (applicable for STS2 and STS4 Cards)
	//
	//1 - Manual
	//
	//2 - Auto (default)
	//
	// +kubebuilder:default:=2
	// +kubebuilder:validation:Required
	SynceRecClkMode int `json:"synceRecClkMode,omitempty"`

	// Configures the synchronization network
	//
	//1 - Option 1 refers to synchronization networks designed for Europe
	//
	//2 - Option 2 refers to synchronization networks designed for United States
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2
	// +kubebuilder:validation:Required
	SyncOption int `json:"syncOption,omitempty"`

	// +kubebuilder:default:=4
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=10
	// Set CPU Pin for SyncE ESMC thread (cpu affinity on each node)
	// +kubebuilder:validation:Required
	StsCpu int `json:"stsCpu,omitempty"`

	//
	// Enable/disable Frame Phase Offset for GM
	// Valid range 0-1
	//
	//     0 - Disable Offset (default)
	//
	//     1 - Enable Offset
	//
	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	PhaseOfstEn_GM int `json:"phaseOfstEn_GM,omitempty"`

	//
	// Set Frame Phase Offset Value for GM
	// Valid range 1947460-1950460 nsec
	//
	// 1948960 - (default)
	//
	// +kubebuilder:default:=1948960
	// +kubebuilder:validation:Minimum=1947460
	// +kubebuilder:validation:Maximum=1950460
	// +kubebuilder:validation:Required
	PhaseOfstVal_GM int `json:"phaseOfstVal_GM,omitempty"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	//Enable/disable two-step PTP Clock
	//
	//0 - Disable two-step clock, (set one-step clock) (default)
	//
	//1 - Enable two-step clock
	// +kubebuilder:validation:Required
	TwoStep int `json:"twoStep,omitempty"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	//Configures the Squelch Mode
	//
	//0 - Sending Sync packet in GM and BC modes while in Freerun (default)
	//
	//1 - Not sending Sync packet in GM and BC modes while in Freerun
	FreerunSQ int `json:"freerunSQ,omitempty"`

	//Set Priority 2 for GM PTP Clock
	//
	// Valid range 0-255, smaller values indicate higher priority
	// +kubebuilder:default:=128
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=255
	// +kubebuilder:validation:Required
	Priority2 int `json:"priority2,omitempty"`

	//Forwardable/Non-Forwardable Multicast Address
	//
	//0 - Non-Forwardable
	//
	//1 - Forwardable (default)
	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// +kubebuilder:validation:Required
	Forwardable int `json:"forwardable,omitempty"`

	//Trace PTP Message
	//
	//Valid range -1-10
	//
	//-1 - Disable Trace log for PTP Messages (default)
	//
	//0 - Enable Trace for all types of PTP Messages
	//
	//1 - Enable Trace for SYNC Messages
	//
	//2 - Enable Trace for DELAY_REQ Messages
	//
	//3 - Enable Trace for PEER_DELAY_REQ Messages
	//
	//4 - Enable Trace for PEER_DELAY_RESP Messages
	//
	//5 - Enable Trace for FOLLOW_UP Messages
	//
	//6 - Enable Trace for DELAY_RESP Messages
	//
	//7 - Enable Trace for PEER_DELAY_FOLLOW_UP Messages
	//
	//8 - Enable Trace for ANNOUNCE Messages
	//
	//9 - Enable Trace for SIGNAL Messages
	//
	//10 - Enable Trace for MANAGEMENT Messages
	// +kubebuilder:default:=-1
	// +kubebuilder:validation:Minimum=-1
	// +kubebuilder:validation:Maximum=10
	// +kubebuilder:validation:Required
	TracePtpMsg int `json:"tracePtpMsg,omitempty"`

	// +kubebuilder:default:=23
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=36
	//Trace module
	//
	//Valid range 0-36
	//
	//0 - Read/Write
	//
	//1 - Init
	//
	//2 - Lan
	//
	//3 - Lan Stats
	//
	//4 - Device specific interrupt
	//
	//5 - System interrupt
	//
	//6 - TS Engine interrupt
	//
	//7 - Packet interrupt
	//
	// 8 - PLL interrupt
	//
	// 9 - Signal Handler
	//
	//10 - TS Packet Stream interrupt
	//
	//11 - Transport Layer interrupt
	//
	//12 - PTP Timestamp interrupt
	//
	//13 - Packet Schedule interrupt
	//
	//14 - Main PTP Engine
	//
	//15 - PTP Best-Master-Clock related
	//
	//16 - PTP Unicast Negotiation related
	//
	//17 - PTP Unicast Discovery related
	//
	//18 - PTP Clock, Port or Stream State related
	//
	//19 - TS RECORD MGR
	//
	//20 - Socket Layer
	//
	//21 - Clock Switch
	//
	//22 - DCO MGR
	//
	//23 - Track Packet Process (default)
	//
	//24 - TOD Manager
	//
	//25 - TSIF
	//
	//26 - MSGQ
	//
	//27 - FPE
	//
	//28 - PTP Foreign Master Table
	//
	//29 - PTSF
	//
	//30 - Notify
	//
	//31 - Signal Pipe Handler
	//
	//32 - G781
	//
	//33 - PTP Timer
	//
	//34 - PTP Tlv
	//
	//35 - HO Utils
	//
	//36 - TSA
	// +kubebuilder:validation:Required
	TraceModule int `json:"traceModule,omitempty"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=8
	// +kubebuilder:validation:Required
	TraceLevel int `json:"traceLevel,omitempty"`

	// +kubebuilder:default:=2
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2
	//Configures the ESMC Mode
	//
	//1 - Manual
	//
	//2 - Auto (default)
	// +kubebuilder:validation:Required
	EsmcMode int `json:"esmcMode,omitempty"`

	// +kubebuilder:default:=1
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2
	//Configures the SSM Mode
	//
	//1 - SSM Code (default)
	//
	//2 - ESSM Code
	// +kubebuilder:validation:Required
	SsmMode int `json:"ssmMode,omitempty"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=8
	// +kubebuilder:validation:Required
	AprLevel int `json:"aprLevel,omitempty"`

	// +kubebuilder:default:=500
	// +kubebuilder:validation:Minimum=300
	// +kubebuilder:validation:Maximum=1800
	// +kubebuilder:validation:Required
	SynceHoldOff int `json:"synceHoldOff,omitempty"`

	//
	// GPSD Debug Level
	// Valid range 1-5
	//
	//1 - Alert
	//
	//2 - Critical (default)
	//
	//3 - Error
	//
	//4 - Warning
	//
	//5 - Notice
	//
	// +kubebuilder:default:=2
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=5
	// +kubebuilder:validation:Required
	GpsdDbgLevel int `json:"gpsdDbgLevel,omitempty"`

	// +kubebuilder:default:=24
	// +kubebuilder:validation:Minimum=24
	// +kubebuilder:validation:Maximum=43
	// +kubebuilder:validation:Required
	DomainNum_8275_1 int `json:"domainNum_8275_1,omitempty"`

	// +kubebuilder:default:=44
	// +kubebuilder:validation:Minimum=44
	// +kubebuilder:validation:Maximum=63
	// +kubebuilder:validation:Required
	DomainNum_8275_2 int `json:"domainNum_8275_2,omitempty"`

	// +kubebuilder:default:=4
	// +kubebuilder:validation:Minimum=4
	// +kubebuilder:validation:Maximum=23
	// +kubebuilder:validation:Required
	DomainNum_8265_2 int `json:"domainNum_8265_2,omitempty"`

	// +kubebuilder:default:=128
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=255
	//Set Clock Local Priority for BC/TSC PTP profiles
	//Valid range 1-255, smaller values indicate higher priority
	// +kubebuilder:validation:Required
	LocalClockPriority int `json:"localClockPriority,omitempty"`

	// +kubebuilder:default:="0,192.168.1.1,,,"
	// Unicast master table
	// for G.8275.2 Boundary Clock profile series specifies unicast masters
	// for its slave ports. There can be multiple entries per port.
	// Each entry must be in the following format:
	//  <port_number>,<master_ip_address>,[<log_announce_interval>],[<log_sync_interval>],[<log_delayreq_interval>]
	// Note:
	//  <port_number> is zero-based, so 0 corresponds to port1 above.
	//  For UDP/IPv6 ports <master_ip_address> shall specify IPv6 address.
	//  Omitting any of the intervals will make that interval take on the default for the profile.
	// Example:
	//  ipv6PortMask = 0x002
	//  unicastMaster = 0,192.168.1.100
	//  unicastMaster = 0,192.168.1.101,,0,0
	//  unicastMaster = 1,fe80::211:22ff:fe33:4455,0,-1,-1
	// +kubebuilder:validation:Required
	UnicastMaster string `json:"unicastMaster,omitempty"`
}

type StsInterfaceSpec struct {
	EthName string `json:"ethName"`

	// +kubebuilder:validation:Minimum=1
	// This is 1 based
	EthPort int `json:"ethPort"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	SyncE int `json:"synce,omitempty"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	Ipv6 int `json:"ipv6,omitempty"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	Ipv4 int `json:"ipv4,omitempty"`

	// +kubebuilder:default:=500
	// +kubebuilder:validation:Minimum=300
	// +kubebuilder:validation:Maximum=1800
	HoldOff int `json:"holdoff,omitempty"`

	// +kubebuilder:validation:Enum=Master;Slave
	// +kubebuilder:default:=Master
	Mode string `json:"mode,omitempty"`

	// +kubebuilder:default:=0
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1
	// Configures QL enable for the interface
	// 0 - Disable QL
	// 1 - Enable QL (default)
	QlEnable int `json:"qlEnable,omitempty"`

	// +kubebuilder:default:=4
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=7
	//Configures QL value for the interface
	//
	//if syncOption is 1 (Europe)
	//
	//===========================
	//
	//1 - QL-PRC
	//
	//2 - QL-PRTC
	//
	//3 - QL-EEC1
	//
	//4 - QL-DNU (default)
	//
	//if syncOption is 2 (United States)
	//
	//==================================
	//
	//5 - QL-PRS
	//
	//2 - QL-PRTC
	//
	//6 - QL-EEC2
	//
	//7 - QL-DUS (default)
	Ql int `json:"ql,omitempty"`

	// +kubebuilder:default:=10000
	// +kubebuilder:validation:Minimum=10000
	// +kubebuilder:validation:Maximum=25000
	//
	//
	PortSpeed int `json:"portSpeed,omitempty"`

	// +kubebuilder:default:=128
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=255
	//
	//Set Port Local Priorities for BC/TSC PTP profiles
	//Valid range 1-255, smaller values indicate higher priority
	//
	LocalPortPriority int `json:"localPortPriority,omitempty"`
}

// StsConfigStatus defines the observed state of StsConfig
type StsConfigStatus struct {
	Nodes []string `json:"nodes,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// StsConfig is the Schema for the stsconfigs API
type StsConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:validation:Required
	Spec   StsConfigSpec   `json:"spec,omitempty"`
	Status StsConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// StsConfigList contains a list of StsConfig
type StsConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StsConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StsConfig{}, &StsConfigList{})
}

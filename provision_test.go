package network_test

import (
	"fmt"
	"testing"

	"github.com/gredinger/network"
)


var sessionStart = `
{{BANNER}}
  End of keyboard-interactive prompts from server
{{HOSTNAME}}>
`

//borrowed from work, anonymized
var shMacAddressOut = `
          Mac Address Table
-------------------------------------------

Vlan    Mac Address       Type        Ports
----    -----------       --------    -----
 All    0305.1ccc.2ccc    STATIC      CPU
 All    0305.1ccc.2ccd    STATIC      CPU
 All    0385.2205.2bab    STATIC      CPU
 All    0385.2205.2b03    STATIC      CPU
 All    0385.2205.2b02    STATIC      CPU
 All    0385.2205.2b03    STATIC      CPU
 All    0385.2205.2b04    STATIC      CPU
 All    0385.2205.2b05    STATIC      CPU
 All    0385.2205.2b06    STATIC      CPU
 All    0385.2205.2b07    STATIC      CPU
 All    0385.2205.2b08    STATIC      CPU
 All    0385.2205.2b09    STATIC      CPU
 All    0385.2205.2b0a    STATIC      CPU
 All    0385.2205.2b0b    STATIC      CPU
 All    0385.2205.2b0c    STATIC      CPU
 All    0385.2205.2b0d    STATIC      CPU
 All    0385.2205.2b0e    STATIC      CPU
 All    0385.2205.2b0f    STATIC      CPU
 All    0385.2205.2b10    STATIC      CPU
 All    ffff.ffff.ffff    STATIC      CPU
   1    ab08.1b15.29a3    DYNAMIC     Gi1/1/26
   1    ab0a.1c49.1a0c    DYNAMIC     Gi1/1/1
   1    ab0f.2ed2.240a    DYNAMIC     Gi1/1/1
   1    ab18.2e7b.12f4    DYNAMIC     Gi1/1/4
   1    ab18.2eab.2a4a    DYNAMIC     Gi1/1/1
   1    ab23.14ba.114e    DYNAMIC     Gi1/1/1
   1    ab23.14e9.2c0d    DYNAMIC     Gi1/1/1
   1    ab23.14e9.2c8d    DYNAMIC     Gi1/1/1
   1    ab24.1b42.1109    DYNAMIC     Gi1/1/1
   1    ab25.11a4.21f7    DYNAMIC     Gi1/1/1
   1    ab26.1305.2fb5    DYNAMIC     Gi1/1/1
   1    ab26.133f.141b    DYNAMIC     Gi1/1/1
   1    ab26.133f.1448    DYNAMIC     Gi1/1/1
   1    ab26.134a.1b67    DYNAMIC     Gi1/1/1
   1    ab26.137c.2410    DYNAMIC     Gi1/1/1
   1    ab26.138b.1517    DYNAMIC     Gi1/1/12
   1    ab26.13a4.178e    DYNAMIC     Gi1/1/1
   1    ab26.13a4.17c4    DYNAMIC     Gi1/1/14
   1    ab26.13d4.228f    DYNAMIC     Gi1/1/1
   1    ab26.13d4.2823    DYNAMIC     Gi1/1/1
   1    ab26.13d4.28df    DYNAMIC     Gi1/1/30
   1    ab26.13e6.2a5f    DYNAMIC     Gi1/1/1
   1    ab35.18b1.2856    DYNAMIC     Gi1/1/1
   1    ab85.120b.13dc    DYNAMIC     Gi1/1/1
   1    ab85.120b.172c    DYNAMIC     Gi1/1/1
   1    ab85.239d.1bfc    DYNAMIC     Gi1/1/1
   1    ab85.23a1.1fcc    DYNAMIC     Gi1/1/1
   1    aba5.1d03.2b89    DYNAMIC     Gi1/1/1
   1    abaf.1fbd.18d8    DYNAMIC     Gi1/1/1
   1    abd8.111d.1924    DYNAMIC     Gi1/1/1
   1    abd8.116a.1e2f    DYNAMIC     Gi1/1/42
   1    abd8.11f1.2eac    DYNAMIC     Gi1/1/44
   1    abe5.1c68.2bdc    DYNAMIC     Gi1/1/1
   1    abe5.2922.12fd    DYNAMIC     Gi1/1/1
   1    047b.2b5a.1b6f    DYNAMIC     Gi1/1/27
   1    047b.2b5b.2a60    DYNAMIC     Gi1/1/40
   1    0805.176a.2c19    DYNAMIC     Gi1/1/1
   1    0805.178b.2b3b    DYNAMIC     Gi1/1/1
   1    0805.17ba.1de2    DYNAMIC     Gi1/1/1
   1    083a.185a.104b    DYNAMIC     Gi1/1/1
   1    0894.2f4f.137c    DYNAMIC     Gi1/1/1
   1    0cd9.16ff.1ec0    DYNAMIC     Gi1/1/1
   1    1062.25b7.2794    DYNAMIC     Gi1/1/1
   1    2c66.1d8b.2ba5    DYNAMIC     Gi1/1/1
   1    2c66.1d8d.2be9    DYNAMIC     Gi1/1/1
   1    2ca5.2878.12c7    DYNAMIC     Gi1/1/1
   1    2ca5.2878.26c5    DYNAMIC     Gi1/1/1
   1    2829.1624.2154    DYNAMIC     Gi1/1/1
   1    2829.1624.2225    DYNAMIC     Gi1/1/1
   1    2cf5.1d18.1fe6    DYNAMIC     Gi1/1/36
   1    2cf5.1d1b.211a    DYNAMIC     Gi1/1/1
   1    2cf5.1d52.2fe3    DYNAMIC     Gi1/1/1
   1    2cf5.1dc1.12a8    DYNAMIC     Gi1/1/1
   1    2cf5.1dc1.12d0    DYNAMIC     Gi1/1/1
   1    2cf5.1dca.2e05    DYNAMIC     Gi1/1/1
   1    2cf5.1dca.2f06    DYNAMIC     Gi1/1/1
   1    2cf5.1dca.2f5c    DYNAMIC     Gi1/1/1
   1    2cf5.1dca.24e6    DYNAMIC     Gi1/1/1
   1    2cf5.1dca.24fd    DYNAMIC     Gi1/1/1
   1    3ab5.1c1b.2430    DYNAMIC     Gi1/1/1
   1    3ab5.1ce4.2bb2    DYNAMIC     Gi1/1/1
   1    309c.15f3.14b3    DYNAMIC     Gi1/1/1
   1    3445.258a.12ce    DYNAMIC     Gi1/1/1
   1    34bd.285e.1bc0    DYNAMIC     Gi1/1/1
   1    34bd.285e.18c0    DYNAMIC     Gi1/1/1
   1    34bd.286e.1d40    DYNAMIC     Gi1/1/1
   1    34bd.287c.2ac0    DYNAMIC     Gi1/1/1
   1    34bd.287c.2e40    DYNAMIC     Gi1/1/1
   1    3822.224b.16c1    DYNAMIC     Gi1/1/1
   1    3822.22c7.1f89    DYNAMIC     Gi1/1/1
   1    3822.22c7.1f8a    DYNAMIC     Gi1/1/1
   1    38f3.2b85.1d62    DYNAMIC     Gi1/1/1
   1    5f2a.2445.16d4    DYNAMIC     Gi1/1/1
   1    5f2a.2461.1111    DYNAMIC     Gi1/1/1
   1    5f2a.2492.1145    DYNAMIC     Gi1/1/1
   1    5f2a.24d5.1c18    DYNAMIC     Gi1/1/13
   1    5fe1.22c2.2387    DYNAMIC     Gi1/1/35
   1    5fe1.21d1.19c3    DYNAMIC     Gi1/1/43
   1    5fe1.21d1.27be    DYNAMIC     Gi1/1/28
   1    5fe1.21d3.1351    DYNAMIC     Gi1/1/1
   1    40f2.2907.1932    DYNAMIC     Gi1/1/1
   1    4409.283e.17e5    DYNAMIC     Gi1/1/46
   1    4437.26ca.196e    DYNAMIC     Gi1/1/1
   1    4437.26e3.1698    DYNAMIC     Gi1/1/1
   1    448a.1b74.1b6c    DYNAMIC     Gi1/1/1
   1    482a.2336.19ae    DYNAMIC     Gi1/1/1
   1    482a.2337.2172    DYNAMIC     Gi1/1/39
   1    4ca6.1d32.1640    DYNAMIC     Gi1/1/1
   1    4cd9.1f5f.1317    DYNAMIC     Gi1/1/1
   1    4cd9.1f4e.1145    DYNAMIC     Gi1/1/1
   1    4cd9.1f74.205f    DYNAMIC     Gi1/1/1
   1    5486.2c5f.1cc0    DYNAMIC     Gi1/1/1
   1    5838.1936.2b20    DYNAMIC     Gi1/1/1
   1    5838.1938.11f9    DYNAMIC     Gi1/1/1
   1    5838.1941.19d4    DYNAMIC     Gi1/1/1
   1    5cff.15d5.1f9c    DYNAMIC     Gi1/1/23
   1    5cff.15db.1d65    DYNAMIC     Gi1/1/31
   1    6805.2a46.1037    DYNAMIC     Gi1/1/1
   1    6827.1939.2931    DYNAMIC     Gi1/1/1
   1    6c0b.143d.1a69    DYNAMIC     Gi1/1/1
   1    6c4b.1ab8.18fa    DYNAMIC     Gi1/1/1
   1    6c4b.1ab8.2109    DYNAMIC     Gi1/1/1
   1    6c4b.105f.1ab8    DYNAMIC     Gi1/1/1
   1    6c4b.1024.2b34    DYNAMIC     Gi1/1/1
   1    6c4b.104e.145a    DYNAMIC     Gi1/1/1
   1    6c4b.1059.2b03    DYNAMIC     Gi1/1/1
   1    6c4b.1059.2d33    DYNAMIC     Gi1/1/1
   1    6cae.1b5d.265c    DYNAMIC     Gi1/1/1
   1    7c4c.1891.20b7    DYNAMIC     Gi1/1/1
   1    7c4c.18c5.23e2    DYNAMIC     Gi1/1/1
   1    7c4c.18c6.2d68    DYNAMIC     Gi1/1/1
   1    7c4c.18d6.1760    DYNAMIC     Gi1/1/1
   1    7c4c.18d9.288a    DYNAMIC     Gi1/1/1
   1    806c.1b03.1b99    DYNAMIC     Gi1/1/1
   1    80e8.1caf.27d1    DYNAMIC     Gi1/1/29
   1    80e8.1f51.2c40    DYNAMIC     Gi1/1/1
   1    80e8.1f76.14c0    DYNAMIC     Gi1/1/1
   1    84a9.1eac.250a    DYNAMIC     Gi1/1/1
   1    88dc.1648.1440    DYNAMIC     Gi1/1/1
   1    90e2.2a7b.2fbe    DYNAMIC     Gi1/1/1
   1    94c6.1129.2d77    DYNAMIC     Gi1/1/1
   1    94c6.1181.2e6d    DYNAMIC     Gi1/1/1
   1    94c6.1183.1d0b    DYNAMIC     Gi1/1/1
   1    98fa.1b29.1c4c    DYNAMIC     Gi1/1/13
   1    d43d.1e61.1bef    DYNAMIC     Gi1/1/1
   1    d43d.1e61.1c0d    DYNAMIC     Gi1/1/9
   1    d89e.237a.1d16    DYNAMIC     Gi1/1/1
   1    d89e.239e.1ed9    DYNAMIC     Gi1/1/1
   1    d8bb.212c.2d1a    DYNAMIC     Gi1/1/1
   1    d8bb.217a.15e8    DYNAMIC     Gi1/1/1
  10    abe5.1707.13f9    DYNAMIC     Gi1/1/1
  10    2cf4.254d.221f    DYNAMIC     Gi1/1/1
  10    2cf4.254d.226d    DYNAMIC     Gi1/1/39
  10    2cf4.254d.2271    DYNAMIC     Gi1/1/1
  10    2cf4.254e.1cd1    DYNAMIC     Gi1/1/1
  10    2cf4.254e.1cde    DYNAMIC     Gi1/1/1
  10    2cf4.254e.1ce0    DYNAMIC     Gi1/1/23
  10    2cf4.254e.1ce1    DYNAMIC     Gi1/1/1
  10    2cf4.254e.1ce5    DYNAMIC     Gi1/1/1
  10    2cf4.254e.1ce7    DYNAMIC     Gi1/1/40
  10    2cf4.254e.1ce8    DYNAMIC     Gi1/1/1
  10    2cf4.254e.1ceb    DYNAMIC     Gi1/1/1
  10    2cf4.254e.1ced    DYNAMIC     Gi1/1/1
  10    2cf4.254e.1cee    DYNAMIC     Gi1/1/1
  10    2cf4.254e.1cf1    DYNAMIC     Gi1/1/1
  10    2cf4.254e.1cf6    DYNAMIC     Gi1/1/33
  10    2cf4.254e.1cf7    DYNAMIC     Gi1/1/25
  10    2cf4.254e.1d26    DYNAMIC     Gi1/1/41
  10    2cf4.254e.1d2b    DYNAMIC     Gi1/1/43
  10    2cf4.254e.1dda    DYNAMIC     Gi1/1/1
  10    2cf4.254e.1164    DYNAMIC     Gi1/1/1
  10    2cf4.254e.1168    DYNAMIC     Gi1/1/1
  10    2cf4.254e.117b    DYNAMIC     Gi1/1/1
  10    2cf4.254e.1203    DYNAMIC     Gi1/1/1
  10    2cf4.254e.1202    DYNAMIC     Gi1/1/1
  10    2cf4.254e.1204    DYNAMIC     Gi1/1/1
  10    2cf4.254e.120c    DYNAMIC     Gi1/1/1
  10    2cf4.254e.1230    DYNAMIC     Gi1/1/45
  10    2cf4.254e.12ac    DYNAMIC     Gi1/1/1
  10    2cf4.254e.12c7    DYNAMIC     Gi1/1/35
  10    2cf4.254e.160f    DYNAMIC     Gi1/1/1
  10    2cf4.2552.1364    DYNAMIC     Gi1/1/47
  10    2cf4.2558.13ab    DYNAMIC     Gi1/1/27
  10    2cf4.2558.198a    DYNAMIC     Gi1/1/10
  10    2cf4.2558.1a1e    DYNAMIC     Gi1/1/1
  10    5fb1.1b4e.2f3e    DYNAMIC     Gi1/1/8
  10    5fb1.1b4e.26cb    DYNAMIC     Gi1/1/1
  10    509a.1c8b.1d5a    DYNAMIC     Gi1/1/1
  10    7038.2ed5.23e1    DYNAMIC     Gi1/1/15
  10    706b.29c6.2808    DYNAMIC     Gi1/1/1
  10    aba9.2deb.106b    DYNAMIC     Gi1/1/1
  10    aba9.2deb.12cd    DYNAMIC     Gi1/1/37
  10    aba9.2deb.24a3    DYNAMIC     Gi1/1/48
  10    aba9.2dec.1c1a    DYNAMIC     Gi1/1/22
  10    b4b5.1799.1af1    DYNAMIC     Gi1/1/1
  10    b6d5.1402.2bab    DYNAMIC     Gi1/1/1
  10    c81f.2ab9.2406    DYNAMIC     Gi1/1/29
  10    c81f.2ab9.29bc    DYNAMIC     Gi1/1/31
  10    c81f.2ac5.1bb6    DYNAMIC     Gi1/1/44
  10    c81f.2ac5.12b1    DYNAMIC     Gi1/1/42
  10    c81f.2ac5.1823    DYNAMIC     Gi1/1/36
  10    c81f.2ac5.1841    DYNAMIC     Gi1/1/28
  10    c81f.2ac5.1915    DYNAMIC     Gi1/1/32
  15    ab1f.163a.144f    DYNAMIC     Gi1/1/1
  15    ab24.1b2b.2044    DYNAMIC     Gi1/1/1
  15    ab24.1b73.18ca    DYNAMIC     Gi1/1/1
  15    ab26.134a.1d38    DYNAMIC     Gi1/1/1
  15    ab68.2b3d.15d0    DYNAMIC     Gi1/1/41
  15    abd8.111d.193a    DYNAMIC     Gi1/1/1
  15    abd8.111d.1c45    DYNAMIC     Gi1/1/1
  15    abd8.1161.13dc    DYNAMIC     Gi1/1/1
  15    abd8.11f2.2bf9    DYNAMIC     Gi1/1/1
  15    047b.2b16.1d25    DYNAMIC     Gi1/1/1
  15    083a.1856.279e    DYNAMIC     Gi1/1/1
  15    083a.185c.1bd1    DYNAMIC     Gi1/1/32
  15    083a.185d.23ab    DYNAMIC     Gi1/1/1
  15    1865.140a.1546    DYNAMIC     Gi1/1/1
  15    2cf5.1dca.2494    DYNAMIC     Gi1/1/1
  15    38d1.1503.10cc    DYNAMIC     Gi1/1/45
  15    5f2a.2445.20cd    DYNAMIC     Gi1/1/1
  15    5fe1.21b7.1521    DYNAMIC     Gi1/1/1
  15    5fe1.21b7.28b7    DYNAMIC     Gi1/1/1
  15    5fe1.21b8.2d6b    DYNAMIC     Gi1/1/37
  15    5fe1.21bd.2dab    DYNAMIC     Gi1/1/1
  15    5fe1.21d1.14fe    DYNAMIC     Gi1/1/33
  15    5fe1.21d1.259e    DYNAMIC     Gi1/1/45
  15    5fe1.21d3.1137    DYNAMIC     Gi1/1/1
  15    482a.2355.1139    DYNAMIC     Gi1/1/1
  15    482a.2355.16ad    DYNAMIC     Gi1/1/1
  15    482a.2358.2255    DYNAMIC     Gi1/1/1
  15    4ccc.1ac6.2825    DYNAMIC     Gi1/1/34
  15    5cb9.13f3.1039    DYNAMIC     Gi1/1/1
  15    5cff.15c2.2d57    DYNAMIC     Gi1/1/1
  15    5cff.15d3.1ec2    DYNAMIC     Gi1/1/47
  15    5cff.15d6.1b60    DYNAMIC     Gi1/1/1
  15    5cff.15dd.2a70    DYNAMIC     Gi1/1/1
  15    5cff.15dd.2ba0    DYNAMIC     Gi1/1/1
  15    5cff.15de.12a0    DYNAMIC     Gi1/1/1
  15    5cff.15de.14f6    DYNAMIC     Gi1/1/48
  15    6827.1939.2a89    DYNAMIC     Gi1/1/1
  15    6c4b.10e3.2b6e    DYNAMIC     Gi1/1/1
  15    d8bb.21bc.1696    DYNAMIC     Gi1/1/1
  15    dc4a.1e5d.211d    DYNAMIC     Gi1/1/22
  15    f439.196b.28a1    DYNAMIC     Gi1/1/1
  15    fc4d.24d7.1a25    DYNAMIC     Gi1/1/1
  40    abbe.15e9.12cc    DYNAMIC     Gi1/1/2
  40    08cc.2787.2ccf    DYNAMIC     Gi1/1/1
  40    08cc.2787.2f6f    DYNAMIC     Gi1/1/1
  40    286f.1f9b.1f0c    DYNAMIC     Gi1/1/1
  40    4c77.1dff.2b78    DYNAMIC     Gi1/1/3
  40    706b.29c6.2808    DYNAMIC     Gi1/1/1
  50    9c7b.2f5a.2122    DYNAMIC     Gi1/1/1
  50    b6d5.1402.2bab    DYNAMIC     Gi1/1/1
  50    e8eb.1bcd.13a6    DYNAMIC     Gi1/1/1
  `

func TestParseMAT(t *testing.T) {
  sw := network.Switch{}
	sw.ParseMAT(shMacAddressOut)
  ports := sw.Ports
	fmt.Println(ports)
	if len(ports) < 20 {
		t.Fail()
	}
	for _, x := range ports {
		if x.Name == "Gi1/1/1" {
			if len(x.Entries) < 20 {
				t.Fail()
			}
		}
	}
}

# tpshp

A GoLang module implementing the TP-Link Smart-Home Protocol. This library is based on the reverse engineering work done by softScheck and documented [here](https://www.softscheck.com/en/reverse-engineering-tp-link-hs110/#TP-Link%20Smart%20Home%20Protocol).

The library has been tested against HS110 smart WiFi plugs but should work for all devices implementing the protocol.

There's also a small command line utility to directly execute TP-Link Smart-Home Protocol commands.
```bash
$ go install github.com/ppacher/tplink-smart-home-protocol/cmd/tpshp
$ tpshp -i 10.8.1.103 system.get_sysinfo {} | jq .
{
  "system": {
    "get_sysinfo": {
      "active_mode": "none",
      "alias": "test",
      "dev_name": "Smart Wi-Fi Plug With Energy Monitoring",
      "deviceId": "80065B1C5FD1C2F230BBD9CCC9DCCE361B1AA47D",
      "err_code": 0,
      "feature": "TIM:ENE",
      "fwId": "00000000000000000000000000000000",
      "hwId": "044A516EE63C875F9458DA25C2CCC5A0",
      "hw_ver": "2.0",
      "icon_hash": "",
      "latitude_i": 489182,
      "led_off": 0,
      "longitude_i": 153172,
      "mac": "D8:0D:17:AD:D3:04",
      "model": "HS110(EU)",
      "next_action": {
        "type": -1
      },
      "oemId": "1998A14DAA86E4E001FD7CAF42868B5E",
      "on_time": 4881,
      "relay_state": 1,
      "rssi": -55,
      "sw_ver": "1.5.4 Build 180815 Rel.121440",
      "type": "IOT.SMARTPLUGSWITCH",
      "updating": 0
    }
  }
}

```
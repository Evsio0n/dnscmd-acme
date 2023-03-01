# Introduction
`dnscmd-acme` is to using `dnscmd` to obtain dns-01 challenge certificate  together with `acme.sh`.

- Warning: This project has ABSOLUTELY NO WARRANTY. Use at your own risk.

## Usage: 

1. Build `acme.exe` and run it on your dns server or change source code to control Active Directory server accrodingly.

```bash
GOOS=windows GOARCH=amd64 go build -o acme.exe
```

2. Copy dns_win.sh to your acme.sh `dns_api` directory.

3. Try `acme.sh` with `--dns dns_win` to obtain certificate.

```bash
acme.sh --issue -d *.example.com -d example.com --dns dns_win --debug 2 
```

## License

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.

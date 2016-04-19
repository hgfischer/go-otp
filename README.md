[![Build Status](https://drone.io/github.com/hgfischer/go-otp/status.png)](https://drone.io/github.com/hgfischer/go-otp/latest)

# OTP

[![Join the chat at https://gitter.im/hgfischer/go-otp](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/hgfischer/go-otp?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Package go-otp implements one-time-password generators used in 2-factor authentication systems like RSA-tokens and Google Authenticator. Currently this supports both HOTP (RFC-4226) and TOTP (RFC-6238).



## Install

```
$ go get github.com/hgfischer/go-otp
```


## Usage

Check API docs, with examples, at http://godoc.org/github.com/hgfischer/go-otp



# License

Copyright (c) 2014, Herbert G. Fischer
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:
    * Redistributions of source code must retain the above copyright
      notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above copyright
      notice, this list of conditions and the following disclaimer in the
      documentation and/or other materials provided with the distribution.
    * Neither the name of the organization nor the
      names of its contributors may be used to endorse or promote products
      derived from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL HERBERT G. FISCHER BE LIABLE FOR ANY
DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

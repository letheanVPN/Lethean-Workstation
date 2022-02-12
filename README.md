# Lethean Workstation

**WORK IN PROGRESS**

[![Build](https://github.com/letheanVPN/workstation/actions/workflows/build.yml/badge.svg)](https://github.com/letheanVPN/workstation/actions/workflows/build.yml)

Lethean Workstation is a remake of a [long decommissioned](http://whois.sc/workstationcommerce.co.uk) system that provided warehousing and distribution as a service in 2010 to a handful of Devon businesses and one in new york; while far from perfect, it was years ahead of Shopify and in heavy development from a warehouse next to Willies Cacao, the chocolate artisan in Devon, England.

The software Workstation was created by members of Letheans new team in 2008; for the original owner of Original Organics Limited, after reading the book [Dreaming in Code](https://en.wikipedia.org/wiki/Dreaming_in_Code), to save the dynamically changing requirements of [Original Organics](http://www.originalorganics.co.uk) as the founder fought to save his company from the financial stock crash of 2008.

It was a parallel processing multi-channel commerce & logistics SaaS application; written in PHP5.

 Born from requirements to survive, not needs, the multi-tenanted application featured, dynamic web interface, before Twig using [PHP_Flexy](https://pear.php.net/manual/en/package.html.html-template-flexy.php), before SNI and implemented Googles COMIT standard to perform cloud print with a bottom-up configuration matrix, dynamically immutable key depths that honoured top-down permission rulesets.

Workstation was written in PHP5, Featured:
- Observable event loop(5) 
- GearMan IPC
- Long Polling Web Workers using COMIT (The original spec)
- Child Companies
- Data environment aware analytical reporting
- PCI Compliant vanity checkout payment page(1)
- Direct integration into: `Parcel Force`, `Home Delivery Network`, `DPD`, `Night Freight`, `Sage`, `Rack Cloud DNS(2)`, `Rackspace Cloud Monitor(2)`, `PHP Notify My Android(3)`
- PDF Generator, able to render any view with a headless browser(4)


 1) old spec did not include the payment page as part of the CDE, don't judge with modern standards; Web1 did not give over to Web2 easily.
 2) One of us authored the PHP OS libs for [Rackspace Cloud DNS](https://github.com/Snider/php-cloudMonitoring), [Rackspace Cloud Monitor](https://github.com/Snider/php-cloudMonitoring) 
 3) One of us authored the [NMA library for PHP5](https://github.com/Snider/php-notifyMyAndroid); linked from NMA as the php5 lib, its use of streams removing external dependency on Curl was not a standard practice.
 4) [HtmlUnit](https://htmlunit.sourceforge.io/).
 5) Yes, PHP, `hook_*_pre`, `hook_*_post`, `mysql`, it worked quite well.


The platform crafted by Lethean's new Lead Developer; remains alive inside their mind.

With the advent of the WebView2 API, Confidential blockchain technology, and networking freedoms we purpose to create, this repositories plan is to bring Workstation back from Web1, where it struggled, to Web3, a place it was designed for, never has data security been so essential to secure business operations in the new digital age, yet to coalesce into reality.

For to long commercial and residential played separately, security needs to be a fundamental right, period. 

Watch this Space for Server-less and DSS compliant Web3 commerce software, business? Remote workers? Regulated space? No problemo. 

Every CryptoNote transaction in every block has its separate ViewKey enabling SOX & HIPPA Compliance by network and deployment configuration by simply saving the keykey in your keyserver DE.

Why should People and Bussiness operate on such alienating terms or have different grades of software functionality?

## Technology Stack

### Window Server

#### Wails.io [Project Information](https://wails.io), [GitHub](https://github.com/wailsapp)

Mid 2021, our lead developer found a project that let people make client native apps; in [Go](https://go.dev/) & HTML/JS/CSS; Wails in v1, was not flexible enough for our needs. But with version two well underway, the project lead sponsored from own funds until the day Lethean could use V2, after the painful experiences with Electron that led to the birth of a [clientside server platform](https://github.com/letheanVPN/server) 

Directly integrated with the operating system [WebView2 API](https://docs.microsoft.com/en-us/microsoft-edge/webview2/) enables PWA's to get drop-in native build's; these can be tied into your CI pipeline and added to your PWA's Manifest file by setting [related_applications](https://developer.mozilla.org/en-US/docs/Web/Manifest/related_applications) and [prefer_related_applications](https://developer.mozilla.org/en-US/docs/Web/Manifest/prefer_related_applications)

#### Client Side Server

Our dVPN is written Typescript using a Rust & Typescript runtime called [Deno](https://deno.land/), the next version of Node.js.

With Deno's impressive set of features, our dVPN executable also ships as an ES Module that utilises WebStandard API's for a seamless development environment that has deliverables for all target platforms.  

When not run from a browser environment, the dVPN Server gives first-class citizenship for the PWA's as a dedicated application installed on their customer's machine.
 
A user creation window on the dVPN's first run instructs users on making what is happening while we take a username and password and create an OpenPGP armoured keypair.

This keypair is the only route to decrypt a configuration object that stores seed information to decrypt the data store; without it knowing the username and password is not enough.

Personal information is protected at all times; both the user name and password are never stored or transmitted; the application then discards the data used once.

In place of a username, a quasi-entropy username is made with a Salt created using a lite l337 conversion and the function input.

Our Zero-Knowledge User Account system is needed to protect journalists and the residents in the field with secure coms back to corporate IT infrastructure; it is that simple.

Knowledge should not be a form of control.

Performing authentication, Two things must be known, the user name and the password to the armoured OpenPGP private key and to physically have that key.

At max privacy settings, data resides within an encrypted namespace that implements the WebStorage API, making it a drop-in replacement for localStorage, sessionStorage, files saved to the Lethean secure namespace are encrypted with the public keys;  identified by the quasi-entropy username.

Integration of external software is made a breeze via multiple integration points.
Cors Enabled Json POST HTTP/TLS Restful Interface.
ZeroMQ WebSockets for IPC
Cors Enabled HTTPS /json_rpc passthrough.
ZeroMQ WebSocket ProcessManager stdOut, stdIn, stdErr

Other Feature include:
	- ProcessManager for executable interactions
 -	namespaced ObjectStore
 -	templated config file server
 -	HTTPS Get Help Documentation
 -	Automatic SDK Builder in over 40 frameworks and languages
        



[![Windows](https://github.com/letheanVPN/lethean/actions/workflows/windows.yml/badge.svg)](https://github.com/letheanVPN/lethean/actions/workflows/windows.yml)
[![MacOS](https://github.com/letheanVPN/lethean/actions/workflows/macos.yml/badge.svg)](https://github.com/letheanVPN/lethean/actions/workflows/macos.yml)

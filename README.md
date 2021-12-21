# Lethean Workstation

**WORK IN PROGRESS**

Lethean Workstation is a remake of a long decommissioned system that provided warehousing and distribution as a service in 2010 to a handful of Devon businesses and one in new york; while far from perfect, it was years ahead of Shopify and in heavy development from a warehouse next to Willies Cacao, the chocolate artisan in Devon, England.

The software Workstation was created by members of Letheans new team in 2008; for the original owner of Original Organics Limited, after reading the book [Dreaming in Code](https://en.wikipedia.org/wiki/Dreaming_in_Code), to save the dynamically changing requirements of [Original Organics](http://www.originalorganics.co.uk) as the founder fought to save his company from the financial stock crash of 2008.

It was a parallel processing multi-channel commerce & logistics SaaS application; written in PHP5.

 Born from requirements to survive, not needs, the multi-tenanted application featured, dynamic web interface, before Twig using [PHP_Flexy](https://pear.php.net/manual/en/package.html.html-template-flexy.php), before SNI and implemented Googles COMIT standard to perform cloud print with a bottom-up configuration matrix, dynamically immutable key depths that honoured top-down permission rulesets.

Workstation was written in PHP5, Featured:
- Observable event loop 
- GearMan IPC
- Long Polling Web Workers using COMIT (The original spec)
- Child Companies
- Data environment aware analytical reporting
- PCI Compliant vanity checkout payment page(1)
- Direct integration into: `Parcel Force`, `Home Delivery Network`, `DPD`, `Night Freight`, `Sage`, `Rack Cloud DNS(2)`, `Rackspace Cloud Monitor(2)`, `PHP Notify My Android(3)`
- PDF Generator, able to render any view with a headless browser(4)


 1) old spec did not include the payment page as part of the CDE, don't judge with modern standards; Web1 did not give over to Web2 easily.
 2) One of us authored the PHP OS libs for [Rackspace Cloud DNS](https://github.com/Snider/php-cloudMonitoring), [Rackspace Cloud Monitor](https://github.com/Snider/php-cloudMonitoring) 
 3) One of us authored the PHP5 NMA library for PHP5; linked from NMA as the php5 lib, its use of streams removing external dependency on Curl was not a standard practice.
 4) [HtmlUnit](https://htmlunit.sourceforge.io/).


Workstation platform crafted by Lethean's new Lead Developer; remains alive inside their mind.

With the advent of the WebView2 API, Confidential blockchain technology, and networking freedoms we purpose to create, this repositories plan is to bring Workstation back from Web1, where it struggled to Web3, where its never been so essential to secure business operations in the new digital age, yet to coalesce into reality.


Watch this Space for Server-less and DSS compliant Web3 commerce software, business? Remote workers? Regulated space? No problemo. 

Every CryptoNote transaction in every block has its separate ViewKey enabling SOX & HIPPA Compliance by network and deployment configuration by simply saving the keykey in your keyserver DE.

Why should Home users, People and Bussiness operate on such alienating terms or have different grades of software functionality?

## Technology Stack

### Window Server

#### Wails.io [Project Information](https://wails.io), [GitHub](https://github.com/wailsapp)

Mid 2021, our lead developer found a project that let people make client native apps; in [Go](https://go.dev/) & HTML/JS/CSS; Wails in v1, was not flexible enough for our needs. But with version two well underway, the project lead sponsored from own funds until the day Lethean could use V2, after the painful experiences with Electron that led to the birth of a [clientside server platform](https://github.com/letheanVPN/server) 

Directly integrated with the operating system [WebView2 API](https://docs.microsoft.com/en-us/microsoft-edge/webview2/) enables PWA's to get drop-in native build's; these can be tied into your CI pipeline and added to your PWA's Manifest file by setting [related_applications](https://developer.mozilla.org/en-US/docs/Web/Manifest/related_applications) and [prefer_related_applications](https://developer.mozilla.org/en-US/docs/Web/Manifest/prefer_related_applications)




[![Windows](https://github.com/letheanVPN/lethean/actions/workflows/windows.yml/badge.svg)](https://github.com/letheanVPN/lethean/actions/workflows/windows.yml)
[![MacOS](https://github.com/letheanVPN/lethean/actions/workflows/macos.yml/badge.svg)](https://github.com/letheanVPN/lethean/actions/workflows/macos.yml)

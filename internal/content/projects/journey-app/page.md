# JourneyApp.me

- Route: `/projects/journey-app`
- Canonical URL: `https://www.winapps.io/projects/journey-app`
- Category: `projects`

## Purpose
Generated from the winapps.io content layer (winapps_next/content) by `npm run export:winny`. Do not edit by hand.

## Winny Notes
- Status: shipped; Creator, sole engineer; Feb 2025 – Present.
- Never state dollar figures or revenue numbers.

## Search Keywords
`projects/journey-app`, `Journey App`, `JourneyApp.me`, `Swift`, `SwiftUI`, `UIKit`, `Go`, `MongoDB`, `NGINX`, `AWS S3`, `JSON Web Tokens`, `bcrypt`, `Let's Encrypt`, `Apple Maps`

## Important Links
- https://apps.apple.com/us/app/journeyapp-me/id6741499202
- https://github.com/mrrobotisreal/JourneyAppWeb
- https://www.winapps.io/products/journey-app/privacy-policy
- https://www.winapps.io/products/journey-app/terms-of-use

## Page Content Snapshot

# JourneyApp.me

> An iOS journaling app in Swift and SwiftUI, with a web companion.

- Status: shipped
- Role: Creator, sole engineer
- Dates: Feb 2025 – Present
- Page: https://www.winapps.io/projects/journey-app

Journey App is a private journal for iPhone and iPad, written in Swift with SwiftUI and UIKit, with a web companion. Entries support markdown, up to ten photos, locations and tags, and full-text search with filters.

I designed and built all of it: the Go backend, the security model, the data layer, and the client.

## Role

I was the sole engineer.

- Designed the backend architecture from scratch.
- Wrote the privacy policy and built authentication and storage that follow it.
- Built the iOS UI.
- Wrote the API specifications and documentation.

## Authentication and encryption

The Go server accepts HTTPS only, with a Let's Encrypt certificate, and redirects plain HTTP.

Each account gets a unique API key alongside JSON Web Tokens, so no one can call an endpoint for another user's data. Keys expire after 90 days and rotate, and every key is rate limited against abuse.

Passwords are salted per user and hashed with bcrypt before they are stored.

Photos live in a private, encrypted S3 bucket. Uploads and views both go through pre-signed URLs, so images stay private while remaining reachable from anywhere.

## Journal entries

I wrote a markdown parser for entry text. It covers headers, bulleted lists, inline code, bold, italic, underline and strikethrough, plus a custom syntax for coloured text.

Each entry holds up to ten photos from the camera or the library, shown as thumbnails at the top of the entry and full size on tap.

Locations come from Apple Maps search inside the app. The entry stores the latitude, longitude, and the address, place name or a custom label. Tags organize entries alongside locations.

## Database, load balancing and search

Journals are personal, so data is stored redundantly. MongoDB runs as a replica set across separate nodes, and sharding spreads query load so responses stay fast.

An NGINX load balancer terminates TLS, taking encryption work off the application servers, and routes traffic with the least-connections algorithm.

Search covers entry text, locations, tags and time ranges. Typing is debounced by one second so the app waits for the user to finish before querying.

## Account settings

Users choose fonts and themes, set a daily writing reminder, and change their password. They can also delete all of their data, which removes every entry, image, location, tag and preference from the database.

## Tech

Swift, SwiftUI, UIKit, Go, MongoDB, NGINX, AWS S3, JSON Web Tokens, bcrypt, Let's Encrypt, Apple Maps

## Links

- [Download on the App Store](https://apps.apple.com/us/app/journeyapp-me/id6741499202)
- [Web app source on GitHub](https://github.com/mrrobotisreal/JourneyAppWeb)
- [Privacy policy](https://www.winapps.io/products/journey-app/privacy-policy)
- [Terms of use](https://www.winapps.io/products/journey-app/terms-of-use)

## Media

- image: Journey App on iPhone: Login — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img970_sm.png
- image: Journey App on iPhone: Signup — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img971_sm.png
- image: Journey App on iPhone: Home 1 — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img940_sm.png
- image: Journey App on iPhone: Home 2 — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img941_sm.png
- image: Journey App on iPhone: Home — Refresh — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img942_sm.png
- image: Journey App on iPhone: Filters — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img943_sm.png
- image: Journey App on iPhone: Filters — Custom Timeframe 1 — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img944_sm.png
- image: Journey App on iPhone: Filters — Custom Timeframe 2 — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img945_sm.png
- image: Journey App on iPhone: Filters — Sort rule, Locations, Tags — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img946_sm.png
- image: Journey App on iPhone: Search — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img972_sm.png
- image: Journey App on iPhone: View Entry — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img947_sm.png
- image: Journey App on iPhone: View Entry — View Image — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img948_sm.png
- image: Journey App on iPhone: View Entry — Tag Details — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img949_sm.png
- image: Journey App on iPhone: View Entry — Location Details — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img950_sm.png
- image: Journey App on iPhone: Edit Entry — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img952_sm.png
- image: Journey App on iPhone: Edit Entry — Scrolled Down — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img953_sm.png
- image: Journey App on iPhone: Edit Entry — Preview — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img954_sm.png
- image: Journey App on iPhone: Create Entry — Empty — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img955_sm.png
- image: Journey App on iPhone: Create Entry — Add Location — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img956_sm.png
- image: Journey App on iPhone: Create Entry — Add Location - Search Places — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img958_sm.png
- image: Journey App on iPhone: Create Entry — Add Tag — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img961_sm.png
- image: Journey App on iPhone: Create Entry — Markdown Help 1 — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img965_sm.png
- image: Journey App on iPhone: Create Entry — Markdown Help 2 — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img966_sm.png
- image: Journey App on iPhone: Create Entry — Markdown Help 3 — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img967_sm.png
- image: Journey App on iPhone: Create Entry — Markdown Help 4 — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img968_sm.png
- image: Journey App on iPhone: Create Entry — Markdown Help 5 — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/img969_sm.png
- image: Journey App on iPad: Login — https://pub-78793ece177d4113887533c21c99d3b4.r2.dev/JourneyApp_iPad001.png
- image: Journey App on iPad: Home — https://pub-78793ece177d4113887533c21c99d3b4.r2.dev/JourneyApp_iPad002.png
- image: Journey App on iPad: Filters — https://pub-78793ece177d4113887533c21c99d3b4.r2.dev/JourneyApp_iPad003.png
- image: Journey App on iPad: Search — https://pub-78793ece177d4113887533c21c99d3b4.r2.dev/JourneyApp_iPad004.png
- image: Journey App on iPad: Create Entry — Write — https://pub-78793ece177d4113887533c21c99d3b4.r2.dev/JourneyApp_iPad005.png
- image: Journey App on iPad: Location Picker — https://pub-78793ece177d4113887533c21c99d3b4.r2.dev/JourneyApp_iPad006.png
- image: Journey App on iPad: Create Entry — Preview — https://pub-78793ece177d4113887533c21c99d3b4.r2.dev/JourneyApp_iPad007.png
- image: Journey App on iPad: View Entry — https://pub-78793ece177d4113887533c21c99d3b4.r2.dev/JourneyApp_iPad008.png
- image: Journey App on iPad: View Image — https://pub-78793ece177d4113887533c21c99d3b4.r2.dev/JourneyApp_iPad009.png
- image: Journey App on iPad: Edit Entry — https://pub-78793ece177d4113887533c21c99d3b4.r2.dev/JourneyApp_iPad010.png

# Senior Software Engineer, People Inc.

- Route: `/portfolio/people-inc`
- Canonical URL: `https://www.winapps.io/portfolio/people-inc`
- Category: `portfolio`

## Purpose
Generated from the winapps.io content layer (winapps_next/content) by `npm run export:winny`. Do not edit by hand.

## Winny Notes
- Status: Senior Software Engineer at People Inc., Aug 2025 – Present (Full-time).
- This role is ongoing.
- Impact is qualitative only: never give a dollar figure or percentage for People Inc revenue.

## Search Keywords
`portfolio/people-inc`, `People Inc.`, `People Inc`, `Senior Software Engineer`, `React Native`, `TypeScript`, `iOS`, `Android`, `Google Mobile Ads (GMA)`, `Google Ad Manager (GAM)`, `Nimbus`, `Liftoff Monetize`, `Node.js`, `Firebase`, `AWS`, `GCP`, `Redis`, `People app`, `prebid`, `prebidding`, `Liftoff`, `Vungle`, `GAM`, `GMA`, `ad revenue`, `mute bug`

## Important Links

## Page Content Snapshot

# Senior Software Engineer, People Inc.

> Senior engineer on the People mobile app's ad revenue team: prebidding integrations, the ad configuration backend, and the tooling around them.

- Dates: Aug 2025 – Present
- Employment: Full-time
- Page: https://www.winapps.io/portfolio/people-inc

## Highlights

- Built the Nimbus and Liftoff Monetize (Vungle) prebidding integrations for the People app from the ground up, working directly with both partner teams. The lift in revenue per ad was large enough that leadership called it out.
- Designed and built the backend ad configuration API, aligning several internal teams on the contract. It serves the People app today and is being adopted by other apps in the company.
- Built an in-app ad debugging panel: toggle each ad partner on or off, see where every config value comes from and what it resolved to, and switch ad test mode without a rebuild.
- Fixed a video ad mute bug that had been open for months. Mute state wasn't shared between the app's video scenes and video ads, and iOS behaved differently when background audio such as Spotify was playing. I synced the state across both and corrected the background-audio handling.
- Built developer tooling the whole team uses: a one-command environment setup script, a screen-recording-to-GIF command-line tool for pull requests and docs, and Rozenite devtools plus mobile-mcp so AI agents can drive iOS and Android simulators and devices.
- Backend work in Node.js across Firebase, AWS, GCP and Redis.

## The ad stack

The People app monetizes with Google Mobile Ads and Google Ad Manager as the mediation layer. In front of the GAM auction we run prebidding: partners bid on an impression before the ad request is sent, and the winning bid is passed into GAM as a price signal so the auction clears higher. I built the Nimbus integration and the Liftoff Monetize integration that do this, working directly with each partner's engineering team on SDK behaviour, bid signalling and edge cases on both platforms.

Both integrations shipped to production and produced a lift in revenue per ad that leadership called out explicitly.

## Ad configuration backend

Ad behaviour used to be scattered across client constants and remote flags. I designed and built a backend configuration API that gives the app one typed, versioned source for every ad parameter: units, partner enablement, refresh rules, test mode, and per-platform overrides. Getting there meant aligning several internal teams on the contract and the rollout.

The API now serves the People app and is being adopted by other apps in the company, which is the result I wanted: a platform, not a feature.

## Ad debugging panel

I built an in-app panel for the ad team that shows every active ad partner and lets you toggle each one, displays where every config value came from and what it resolved to, and switches test mode on or off without a rebuild. It turned an hours-long "which config is this device using" question into a glance.

## The mute bug

A video-ad mute bug had been open for months. Two things were wrong. The app's own video scenes and the video ad SDK each kept their own mute state and never agreed, so ads came in loud or silent at random. On iOS, background audio such as Spotify changed the behaviour again because the audio session category was being set inconsistently. I traced both, introduced one shared mute source of truth that both players observe, and fixed the audio-session handling so background music behaves the way users expect.

## Tooling and backend

Beyond the ad work I build internal tools that shorten the team's debug and release loops. The setup script checks a new machine and installs the right Node, CocoaPods, fastlane and Xcode command-line tools, pulls in the app's private config files, places the debug keystore, and runs the full install. When CocoaPods fails it retries with fixes for the most common failure causes, and it ends with a summary of everything it did. So far it has resolved every setup error the team has hit.

A second tool installs itself as a global Bash/Zsh command, checks for and installs ffmpeg and gifsicle, and turns `.mov` screen recordings into animated GIFs for pull requests and documentation, with configurable size, dimensions and quality. It also converts recordings to other formats such as `.mp4`.

I also added Rozenite devtools and mobile-mcp to the People app. Rozenite made day-to-day debugging much easier, and mobile-mcp lets our AI agents drive iOS simulators and devices and Android emulators and devices directly.

On the backend I work in Node.js across Firebase, AWS, GCP and Redis, including the ad configuration service above.

## Tech

React Native, TypeScript, iOS, Android, Google Mobile Ads (GMA), Google Ad Manager (GAM), Nimbus, Liftoff Monetize, Node.js, Firebase, AWS, GCP, Redis

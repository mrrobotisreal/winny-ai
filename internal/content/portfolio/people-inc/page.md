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
`portfolio/people-inc`, `People Inc.`, `People Inc`, `Senior Software Engineer`, `React Native`, `TypeScript`, `iOS`, `Android`, `Google Mobile Ads (GMA)`, `Google Ad Manager (GAM)`, `Nimbus`, `Liftoff Monetize`, `Node.js`, `Firebase`, `AWS`, `GCP`, `Redis`, `People app`, `prebid`, `prebidding`, `Liftoff`, `Vungle`, `GAM`, `GMA`, `ad revenue`, `mute bug`, `ads configuration`, `feature flags`, `app-ads`, `ads-lite`, `AI Cost Explorer`, `setup scripts`

## Important Links

## Page Content Snapshot

# Senior Software Engineer, People Inc.

> Senior engineer on the People mobile app's ad revenue team: prebidding integrations, the ads configuration system, and the tooling around them.

- Dates: Aug 2025 – Present
- Employment: Full-time
- Page: https://www.winapps.io/portfolio/people-inc

## Highlights

- Built the Nimbus and Liftoff Monetize (Vungle) prebidding integrations for the People app from the ground up, working directly with both partner teams. Nimbus prebidding, integrated into the app's Google Ad Manager and Google Mobile Ads stack, significantly increased ad revenue.
- Led the design and architecture of a new ads configuration system. Ad configs and feature flags can be added, removed or changed without a pull request and without restarting anything, and it is app-agnostic, so other app verticals such as MyRecipes can use it.
- Working directly with the backend team, built a much faster and more efficient feature-flag and config parsing system in the People app's backend. It fetches ad configs and flags from the new ads configuration system and gives paid subscribers an ads-lite experience.
- Created the initial app-ads module, hosted in the company's internal package registry so it installs like any npm package and future apps can reuse the same ad settings and integrations.
- Built an in-app ad debugging panel: toggle each ad partner on or off, see where every config value comes from and what it resolved to, and switch ad test mode without a rebuild.
- Fixed a video ad mute bug that had been open for months. Mute state wasn't shared between the app's video scenes and video ads, and iOS behaved differently when background audio such as Spotify was playing. I synced the state across both and corrected the background-audio handling.
- Leading the team's AI agent workflow initiative (CLAUDE.md, AGENTS.md and custom skills built for our repos and apps), and built an AI Cost Explorer and benchmark that tracks exact token spend per user, per agent and per model, with anomaly detection.
- Built developer tooling the whole team uses: setup scripts that take a new engineer from cloning the repo to a running app in about 15–20 minutes instead of days, a screen-recording-to-GIF command-line tool, and Rozenite devtools plus mobile-mcp so AI agents can drive iOS and Android simulators and devices. Backend work in Node.js across Firebase, AWS, GCP and Redis.

## The ad stack

The People app monetizes with Google Mobile Ads and Google Ad Manager as the mediation layer. In front of the GAM auction we run prebidding: partners bid on an impression before the ad request is sent, and the winning bid is passed into GAM as a price signal so the auction clears higher. I built the Nimbus integration and the Liftoff Monetize integration that do this, working directly with each partner's engineering team on SDK behaviour, bid signalling and edge cases on both platforms.

Both integrations shipped to production. Integrating Nimbus's prebidding SDK into the app's Google Ad Manager and Google Mobile Ads stack significantly increased ad revenue, and leadership called out the lift explicitly.

## Ads configuration system

Ad behaviour used to be scattered across client constants and remote flags. I led the design and architecture of a new ads configuration system that gives the app one source for every ad config and feature flag: units, partner enablement, refresh rules, test mode, and per-platform overrides. Configs and flags can be added, removed or changed without a pull request and without restarting the system.

The system is app-agnostic, so other app verticals in the company, such as MyRecipes, can use it too. That was the result I wanted: a platform, not a feature.

I also created the initial app-ads module. It is registered and hosted in the company's internal package registry, so it installs like a regular npm package, and future apps can reuse the same ad settings and integrations instead of rebuilding them.

## Flag and config parsing

Working directly with the backend team, I built a much faster and more efficient feature-flag and config parsing system in the People app's backend. It fetches ad configs and feature flags from the new ads configuration system and parses them for the app.

It also handles the new subscriber system: when the user is a paid subscriber, it alters the parsed ad configs and flags to give them an ads-lite experience.

## Ad debugging panel

I built an in-app panel for the ad team that shows every active ad partner and lets you toggle each one, displays where every config value came from and what it resolved to, and switches test mode on or off without a rebuild. It turned an hours-long "which config is this device using" question into a glance.

## The mute bug

A video-ad mute bug had been open for months. Two things were wrong. The app's own video scenes and the video ad SDK each kept their own mute state and never agreed, so ads came in loud or silent at random. On iOS, background audio such as Spotify changed the behaviour again because the audio session category was being set inconsistently. I traced both, introduced one shared mute source of truth that both players observe, and fixed the audio-session handling so background music behaves the way users expect.

## AI agent workflows

I'm leading the team's initiative to build an automated agent workflow system: CLAUDE.md and AGENTS.md files and custom skills written specifically for the team's repos and apps.

To go with it I built an AI Cost Explorer and benchmark. It tracks and calculates exact token spend (input, output, cache write and cache read) per user and per agent, and the exact cost per model. Anomaly detection flags unusual spend, and insight charts show where most of the token spend goes, which models are used most for which tasks, and where the team can save money while getting the same or better results.

## Tooling and backend

Beyond the ad work I build internal tools that shorten the team's setup, debug and release loops. My setup scripts take a new engineer from cloning the People app repo to a running app in about 15–20 minutes, where it used to take days; most of that time is downloading and installing. They check a new machine and install the right Node, CocoaPods, fastlane and Xcode command-line tools, pull in the app's private config files, place the debug keystore, and run the full install and build. When a common step such as CocoaPods fails, they retry with small adjustments based on the failure until the install and build succeed, and they end with a summary of everything they did.

A second tool installs itself as a global Bash/Zsh command, checks for and installs ffmpeg and gifsicle, and turns `.mov` screen recordings into animated GIFs for pull requests and documentation, with configurable size, dimensions and quality. It also converts recordings to other formats such as `.mp4`.

I also added Rozenite devtools and mobile-mcp to the People app. Rozenite made day-to-day debugging much easier, and mobile-mcp lets our AI agents drive iOS simulators and devices and Android emulators and devices directly.

On the backend I work in Node.js across Firebase, AWS, GCP and Redis, including the flag and config parsing above.

## Tech

React Native, TypeScript, iOS, Android, Google Mobile Ads (GMA), Google Ad Manager (GAM), Nimbus, Liftoff Monetize, Node.js, Firebase, AWS, GCP, Redis

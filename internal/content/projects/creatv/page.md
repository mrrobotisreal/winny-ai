# CreaTV

- Route: `/projects/creatv`
- Canonical URL: `https://www.winapps.io/projects/creatv`
- Category: `projects`

## Purpose
Generated from the winapps.io content layer (winapps_next/content) by `npm run export:winny`. Do not edit by hand.

## Winny Notes
- Status: active; Founder, lead engineer; Nov 2025 – Present.
- CreaTV is public with creators onboarding; auth is moving from Supabase to Firebase.
- Never state dollar figures or revenue numbers.

## Search Keywords
`projects/creatv`, `CreaTV`, `Go`, `TypeScript`, `React`, `React Native`, `Expo`, `Postgres`, `Redis`, `FFmpeg`, `Python`, `RabbitMQ`, `AWS S3`, `CloudFront`, `Cloudflare Stream`, `Prometheus`, `Grafana`, `CreaTV Studio`, `creator platform`, `video platform`, `streaming`, `creatv.io`

## Important Links
- https://www.creatv.io
- https://www.winapps.io/download/creatv-mobile/android

## Page Content Snapshot

# CreaTV

> A creator-first video platform, built as twenty-plus services.

- Status: active
- Role: Founder, lead engineer
- Dates: Nov 2025 – Present
- Page: https://www.winapps.io/projects/creatv

CreaTV is the video platform I wanted to exist: creator-first monetization, serious attention to playback performance and observability, and a media pipeline that scales. I founded it and lead the engineering, from upload and transcoding through streaming, search, recommendations, payments and the mobile app.

Media Manipulator's Darkroom editor is embedded directly in CreaTV Studio, so creators can edit in the browser before they publish.

## In numbers

| Metric | Value |
| --- | --- |
| Backend services | 20+ |
| Clients | web, Android, iOS |

## What it is

CreaTV is a creator-first video platform I founded in late 2025 and lead the engineering for. The goal is simple to state and hard to build: monetization that favours the creator, playback that feels instant, and a media pipeline that does not fall over when it gets popular.

It ships as a web app, an Android app and an iOS app, backed by more than twenty independent services.

## Architecture

The backend is split by responsibility: an API gateway in front of user, upload, transcoding, video-metadata, streaming, search, recommendation, notification, sharing, snippet, payment, ads, support and assistant services, with shared migrations and an analytics service. The services are written in Go, with Postgres for data, Redis for caching and RabbitMQ between services.

Upload goes straight to object storage with presigned URLs; transcoding produces adaptive renditions; streaming serves HLS and DASH. Video segments, thumbnails, captions, storyboards and sprites live in S3 behind CloudFront, with Cloudflare Stream alongside it for playback. Search and recommendations are their own services so they can scale and fail independently of playback.

Auth runs on Supabase today and is moving to Firebase Auth.

- API gateway + 20 services
- Go services, Postgres, Redis, RabbitMQ
- HLS and DASH from S3 + CloudFront, plus Cloudflare Stream

## Studio and the Darkroom embed

Creators edit before they publish without leaving CreaTV. CreaTV Studio embeds Media Manipulator's Darkroom editor in an iframe and talks to it over a versioned, origin-validated postMessage protocol, so the editor can evolve on its own release cadence.

## Mobile

The mobile app is React Native with Expo for Android and iOS. An Android development build is available on this site.

## Operations

The services run directly on my home server, without containers, and media storage and delivery run on AWS. On-demand EC2 workers for transcoding jobs are in progress. Prometheus and Grafana cover metrics and dashboards, and deploys are manual for now. Performance and observability were design goals from the first commit, not retrofits.

## Where it is

CreaTV is public, and creators are onboarding now.

## Tech

Go, TypeScript, React, React Native, Expo, Postgres, Redis, FFmpeg, Python, RabbitMQ, AWS S3, CloudFront, Cloudflare Stream, Prometheus, Grafana

## Links

- [Visit creatv.io](https://www.creatv.io)
- [Android dev build](https://www.winapps.io/download/creatv-mobile/android)

# Media Manipulator

- Route: `/projects/media-manipulator`
- Canonical URL: `https://www.winapps.io/projects/media-manipulator`
- Category: `projects`

## Purpose
Generated from the winapps.io content layer (winapps_next/content) by `npm run export:winny`. Do not edit by hand.

## Winny Notes
- Status: shipped; Founder, sole engineer; May 2025 – Present.
- Never state dollar figures or revenue numbers.

## Search Keywords
`projects/media-manipulator`, `Media Manipulator`, `Go`, `Gin`, `Postgres`, `Redis`, `FFmpeg`, `ImageMagick`, `NVENC`, `Next.js`, `React`, `TypeScript`, `Tailwind`, `WebGL2`, `Web Audio`, `Whisper`, `Real-ESRGAN`, `Demucs`, `Ollama`, `S3`, `Firebase Auth`, `Vercel`, `nginx`, `Cloudflare`, `Darkroom`, `Content Studio`, `video editor`, `media converter`, `AI restoration`, `document scan`

## Important Links
- https://www.media-manipulator.com
- https://www.media-manipulator.com/tools/content-studio

## Page Content Snapshot

# Media Manipulator

> 65+ browser media tools and Darkroom, a multi-track video editor, on GPU servers I run myself.

- Status: shipped
- Role: Founder, sole engineer
- Dates: May 2025 – Present
- Page: https://www.winapps.io/projects/media-manipulator

Media Manipulator converts, compresses, trims, transcribes, restores and edits images, video and audio in the browser, with the heavy work done on my own GPU servers so large files and AI models are no problem. Darkroom, its multi-track timeline editor, brings colour grading, chroma key, keyframes, transitions, titles and captions to a web page and exports through ffmpeg with preview-to-export parity.

The AI tools (restoration, upscaling, transcription, document scanning, face and PII redaction) run on hardware I operate. No user file goes to a third-party AI provider.

## In numbers

| Metric | Value |
| --- | --- |
| Tools | 65+ |
| Locales | 6 |
| Transition types | 15 |
| Blend modes | 8 |
| AI models served | 15+ |
| API routes | ~130 |
| Go tests | 490+ |

## What it is

Media Manipulator started as a converter and became a media workbench. It has 65+ single-purpose tools for images, video and audio, each with its own page, plus Darkroom, a multi-track video editor that runs in a browser tab. Files are processed on servers I run, deleted within 24 hours, and no account is needed for the standard tools.

The site ships in six locales, including a right-to-left one, with a key-parity test so no translation can silently lag. It is a CreaTV Ltd. product, and Darkroom is embedded directly in CreaTV Studio.

## Darkroom: a timeline editor in the browser

Darkroom composites video, image, audio and LUT tracks in a WebGL2 pipeline: crop, fit, transform, colour grade, 3D LUT, chroma key, opacity and blend, in that order, with a CSS-filter fallback if the GPU context is lost. Audio mixes through a Web Audio graph. Decoders are pooled and Safari's decoder cap is respected.

Editing features: trim, split and ripple delete at frame accuracy; a media bin with server-rendered proxies, filmstrips and waveform peaks; a Lumetri-style grade with exposure, contrast, saturation, temperature, tint and vibrance; chroma key with an eyedropper; eight blend modes; fifteen transitions; keyframes on position, scale, rotation, opacity, volume, pan and any effect parameter with five easing curves; title clips with text and shape layers; audio ducking and loudness presets; a captions lane with Whisper auto-captions and SRT/VTT export or burn-in.

Projects are safe to work in: snapshot undo/redo, revision-numbered saves that refuse to clobber, crash-recovery drafts in IndexedDB, and version history with checkpoints. Export is a server-side ffmpeg graph with hardware encoding when available, and a golden-image harness keeps the preview and the export within a documented tolerance.

- Target scale: 60-minute timelines, 12 tracks, 1,000 clips
- Keyframed parameters export through ffmpeg sendcmd
- Embedded in CreaTV Studio over a versioned, origin-checked postMessage protocol

## The tool catalogue

Image: format conversion across JPG, PNG, WebP, HEIC, AVIF, SVG and ICO; resize, compress, background removal, image-to-PDF and PDF-to-image. Video: conversion between MP4, MOV, WebM, AVI and MKV; compress, trim, crop, resize, rotate; audio extraction and replacement; frame extraction; GIF; HLS packaging with an H.264 ladder and DASH packaging with AV1/VP9, multilingual captions, storyboards and an I-frame playlist. Audio: conversion, waveform rendering, vocal isolation.

Privacy: EXIF, GPS, IPTC and XMP removal; face blur; text and PII redaction; object removal.

## AI on hardware I run

The AI tools run on my own GPU server, so no user file is sent to a third-party model provider. Transcription and captions use Whisper; caption translation uses a TranslateGemma model served by Ollama; video restoration offers six models side by side; image restoration combines upscalers with face-restoration and pre-cleaning models; document scanning turns printed pages into a searchable PDF through Tesseract and transcribes handwriting with a vision model that flags uncertain passages instead of guessing. A VRAM-aware scheduler splits jobs across two GPUs.

- Restoration: Real-ESRGAN, SwinIR, HAT, BasicVSR++, RVRT, VRT
- Faces and cleanup: GFPGAN, CodeFormer, FBCNN, SCUNet, NAFNet
- Also: RIFE frame interpolation, Demucs vocal isolation, DeepFilterNet denoise, LaMa inpainting, rembg

## Backend

A Go API (Gin, pgx, Redis) with roughly 130 routes, Postgres with 50+ migrations, presigned S3 uploads for large video, a job manager with polling and server-sent events, tiered quotas with per-route rate limits, Prometheus metrics, and FFmpeg, ImageMagick, ExifTool, Poppler and friends behind it. It runs as a systemd service on a headless Ubuntu workstation behind nginx and Cloudflare, with 490+ Go tests.

## Frontend

Next.js App Router on Vercel, React 19, Tailwind 4, Radix, TanStack Query and Zustand. Six locales served as i18next shards per feature, hreflang alternates in the sitemap, consent-gated first-party analytics, and an ads architecture that is off by default and audited in CI.

## Tech

Go, Gin, Postgres, Redis, FFmpeg, ImageMagick, NVENC, Next.js, React, TypeScript, Tailwind, WebGL2, Web Audio, Whisper, Real-ESRGAN, Demucs, Ollama, S3, Firebase Auth, Vercel, nginx, Cloudflare

## Links

- [Visit media-manipulator.com](https://www.media-manipulator.com)
- [Open Darkroom](https://www.media-manipulator.com/tools/content-studio)

## Media

- image: Media Manipulator home page: the upload panel, quick-tool chips and the Darkroom entry point — https://www.winapps.io/media-manipulator/home.webp
- image: Tool catalogue listing Darkroom first, then image tools such as compress, resize and background removal — https://www.winapps.io/media-manipulator/tools.webp
- image: Darkroom landing page with its quick guide to the multi-track timeline, supported media and export — https://www.winapps.io/media-manipulator/darkroom.webp
- image: AI image restoration tool describing its pre-clean, upscaler and face-enhancement model pipeline — https://www.winapps.io/media-manipulator/restoration.webp
- image: AI document scan tool that turns printed pages and handwriting into a searchable PDF or Word file — https://www.winapps.io/media-manipulator/docscan.webp

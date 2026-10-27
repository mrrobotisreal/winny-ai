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
`projects/media-manipulator`, `Media Manipulator`, `Go`, `Gin`, `Postgres`, `Redis`, `FFmpeg`, `ImageMagick`, `NVENC`, `Next.js`, `React`, `TypeScript`, `Tailwind`, `WebGL2`, `Web Audio`, `Zod`, `WebCodecs`, `mediabunny`, `anime.js`, `Stripe`, `Whisper`, `Real-ESRGAN`, `Demucs`, `Ollama`, `S3`, `Firebase Auth`, `Vercel`, `nginx`, `Cloudflare`, `Darkroom`, `Content Studio`, `video editor`, `media converter`, `AI restoration`, `document scan`

## Important Links
- https://www.media-manipulator.com
- https://www.media-manipulator.com/tools/content-studio

## Page Content Snapshot

# Media Manipulator

> 75+ browser media tools, a multi-track video editor and an SVG animator, on GPU servers I run myself.

- Status: shipped
- Role: Founder, sole engineer
- Dates: May 2025 – Present
- Page: https://www.winapps.io/projects/media-manipulator

Media Manipulator converts, compresses, trims, transcribes, restores and edits images, video and audio in the browser, with the heavy work done on my own GPU servers so large files and AI models are no problem. Darkroom, its multi-track timeline editor, brings colour grading, chroma key, keyframes, transitions, titles and captions to a web page and exports through ffmpeg with preview-to-export parity.

The AI tools (restoration, upscaling, transcription, document scanning, face and PII redaction) run on hardware I operate. No user file goes to a third-party AI provider.

## In numbers

| Metric | Value |
| --- | --- |
| Tool pages | 75+ |
| Locales | 18 |
| Transition types | 15 |
| Metadata tags catalogued | 1,800+ |
| Blend modes | 8 |
| Motion export formats | 11 |
| Motion animators | 6 |
| Pro Motion exports a day | 150 |
| AI models served | 15+ |
| API routes | ~180 |
| Go tests | 900+ |

## What it is

Media Manipulator started as a converter and became a media workbench. It has 75+ single-purpose tools for images, video and audio, each with its own page, plus Darkroom, a multi-track video editor that runs in a browser tab. Files are processed on servers I run. No account is needed for the standard tools, and guests' files are deleted within 24 hours; signed-in users choose their own retention, from 24 hours to 2 years.

Motion, an SVG animation editor in beta, sits beside Darkroom: it turns an SVG into layers on a keyframe timeline and exports animated SVG, Lottie, React, JavaScript, Flutter or GIF straight from the browser.

The site is fully translated into eighteen languages (English, Ukrainian, Russian, Hebrew, German, Spanish for Spain and Latin America, French, Italian, Portuguese for Portugal and Brazil, Japanese, Simplified and Traditional Chinese, Korean, Arabic, Hindi and Vietnamese), two of them right to left, with a key-parity test so no translation can silently lag. It is a CreaTV Ltd. product, and Darkroom is embedded directly in CreaTV Studio.

## Darkroom: a timeline editor in the browser

Darkroom composites video, image, audio and LUT tracks in a WebGL2 pipeline: crop, fit, transform, colour grade, 3D LUT, chroma key, opacity and blend, in that order, with a CSS-filter fallback if the GPU context is lost. Audio mixes through a Web Audio graph. Decoders are pooled and Safari's decoder cap is respected.

Editing features: trim, split and ripple delete at frame accuracy; a media bin with server-rendered proxies (one keyframe a second, so preview seeks stay cheap, and a one-click rebuild for clips imported before that), filmstrips and waveform peaks; a Lumetri-style grade with exposure, contrast, saturation, temperature, tint and vibrance; chroma key with an eyedropper; eight blend modes; fifteen transitions; keyframes on position, scale, rotation, opacity, volume, pan and any effect parameter with five easing curves; title clips with text and shape layers; audio ducking and loudness presets; a captions lane with Whisper auto-captions and SRT/VTT export or burn-in.

Still images can be imported and held on the timeline for any length, Premiere-style, and they export like any other clip. Clips can be sped up or slowed down from 0.1x to 10x with pitch kept, reversed, or frozen on a frame. The timeline has a full keyboard-driven tool set (razor, slip, slide, roll and ripple edits, markers, groups, track lock, solo and hide, and detaching audio from video), and a mixer gives each track a fader, pan, three-band EQ, a compressor and live meters, with a master bus.

The media bin has search, folders, colour labels and rename, a safe delete that lists every clip still using a file, a parallel upload queue, and voiceover recorded straight from the microphone onto the timeline.

Projects are safe to work in: snapshot undo/redo, revision-numbered saves that refuse to clobber, crash-recovery drafts in IndexedDB, and version history with checkpoints. Export is a server-side ffmpeg graph with hardware encoding when available: H.264 or HEVC video, or MP3 and WAV audio, at source resolution or 720p to 2160p, with captions burned in or as SRT/VTT, canvas presets for 16:9, 9:16, 1:1 and 4:5, a frame grab as PNG, and an exports panel with progress and re-render. A golden-image harness keeps the preview and the export within a documented tolerance, and solo and hide apply to the export exactly as they do to the preview.

- Target scale: 60-minute timelines, 12 tracks, 1,000 clips
- Keyframed parameters export through ffmpeg sendcmd
- Embedded in CreaTV Studio over a versioned, origin-checked postMessage protocol, with capability flags and published integration docs
- End-to-end browser tests cover a full edit-save-export journey and crash recovery

## Motion: SVG animation in the browser

Motion is a separate editor from Darkroom, with its own document model: a versioned, additive-only VectorDoc (Zod on the client, a Go sanitizer mirror on the server, shared byte-parity fixtures). An imported SVG becomes a layer tree of groups, paths, shapes, text, images, gradients, masks and clip paths; markup it cannot model is kept as an opaque layer that round-trips unchanged, and anything unsafe (scripts, event handlers, external references, foreignObject) is dropped on import and rebuilt from an allow-list, never inserted as HTML. New artwork comes from rectangle, ellipse, line, pen and text tools.

Every animated property is its own track of keyframes with linear, ease-in, ease-out and ease-in-out presets, a cubic-bezier curve editor and hold keys. Six one-click animators cover stroke draw-on, motion along a path, shape morphing (an in-house normaliser that matches command counts, start points and winding), gradient stops, masks and clip paths, and blur. One sampler is the single definition of a value at time t: the preview renders from it, and every exporter is tested against it.

Export happens in the browser, and the server only counts it against a separate daily Motion allowance before the file is built. Animated SVG is pure CSS when Chrome, Firefox and Safari can all play it, and otherwise embeds a tree-shaken anime.js runtime of about 14 KB gzip. The Lottie writer targets spec 1.0.1 and has a matching reader; a single unsupported-feature list means nothing is dropped silently in either direction. The React, JavaScript/TypeScript and Flutter exporters (a native flutter_svg widget when the animation allows, otherwise a WebView player) are pure functions with golden fixtures. GIF uses a clean-room GIF89a encoder in a Worker, PNG sequences stream into a zip, and MP4 and WebM go through WebCodecs and mediabunny; the video formats are part of the Pro plan.

Licensing is a design constraint, not an afterthought. GSAP is excluded because its free licence forbids no-code animation builders, and Theatre.js Studio is excluded because it is AGPL. A test fails the build if either enters the dependency tree. Every runtime that ships inside an export carries its licence header, and an open-source licences page lists the MPL-2.0 muxer.

- Formats: static SVG (free), animated SVG, Lottie JSON, React, JavaScript, TypeScript, Flutter and GIF, plus MP4, WebM and PNG sequence on Pro (up to 4K, 60 fps, 60 seconds)
- Daily Motion exports by plan: 2 without an account, 5 free, 30 Premium, 150 Pro. The export dialog locks each format a plan lacks and names the plan that adds it; the server stays the authority
- Projects: browser-only drafts in IndexedDB, or account projects with revision-checked saves, autosave and a thumbnail
- Plan table on the tool page rendered from the live tier API, so a retuned allowance needs no deploy

## See the result before you download it

Every tool shows its output in the browser before anything is downloaded, with viewers for images, video, GIF and audio. Compare modes include a wipe slider, side-by-side, synced A/B video with an fps readout, and audio A/B. Transcripts and captions open in viewers with click-to-seek and a two-column translation view; there is also a ZIP archive gallery, a PDF viewer, and HLS and DASH stream playback. For AI video restoration, the original can be compared against each model's output.

## Metadata inspector

The inspector reads EXIF, XMP, IPTC and QuickTime metadata through exiftool and ffprobe and explains it instead of dumping raw tags. It is backed by a catalogue of about 1,800 tags in 22 categories, flags what matters for privacy (location, identity, device, history), plots GPS coordinates on a locator, and shows a before/after diff in split or unified view. "Strip everything" covers video and audio as well as images.

## The tool catalogue

Image: format conversion across JPG, PNG, WebP, HEIC, AVIF, SVG and ICO; resize, compress, background removal, image-to-PDF and PDF-to-image. Video: conversion between MP4, MOV, WebM, AVI and MKV; compress, trim, crop, resize, rotate; audio extraction and replacement; frame extraction; GIF; HLS packaging with an H.264 ladder and DASH packaging with AV1/VP9, multilingual captions, storyboards and an I-frame playlist. Audio: conversion, waveform rendering, vocal isolation, voice cleanup and background-noise removal. Image: upscaling.

Privacy: EXIF, GPS, IPTC and XMP removal; face blur; text and PII redaction; object removal.

Seven of the tools are recent AI additions: upscale, blur faces, redact text, remove object, clean voice, remove background noise and remove vocals.

## AI on hardware I run

The AI tools run on my own GPU server, so no user file is sent to a third-party model provider. Transcription and captions use Whisper; caption translation uses a TranslateGemma model served by Ollama; video restoration offers six models side by side; image restoration combines upscalers with face-restoration and pre-cleaning models; document scanning turns printed pages into a searchable PDF through Tesseract and transcribes handwriting with a vision model that flags uncertain passages instead of guessing. A VRAM-aware scheduler splits jobs across two GPUs.

- Restoration: Real-ESRGAN, SwinIR, HAT, BasicVSR++, RVRT, VRT
- Faces and cleanup: GFPGAN, CodeFormer, FBCNN, SCUNet, NAFNet
- Also: RIFE frame interpolation, Demucs vocal isolation, DeepFilterNet denoise, LaMa inpainting, rembg

## Backend

A Go API (Gin, pgx, Redis) with roughly 180 routes, Postgres with versioned migrations, presigned S3 uploads for large video, a job manager with polling and server-sent events, tiered quotas with per-route rate limits, Stripe billing for the Premium tier, user-selectable data retention with export and delete, transactional email through Resend, a public status page fed by a health endpoint, Prometheus metrics, and FFmpeg, ImageMagick, ExifTool, Poppler and friends behind it. It runs as a systemd service on a headless Ubuntu workstation behind nginx and Cloudflare, with 900+ Go tests.

## Accounts, Premium and Pro

Accounts use Firebase Auth. The account screen has tabs for history, projects, billing, privacy and data, and profile; users can export their data or delete their account, share a result through an expiring, revocable link, and choose which emails they get. Guests get a browser-side Recent history. Premium is 8 USD or 8 EUR a month, or 80 USD or 80 EUR a year, through Stripe, and brings larger files, more jobs and exports, output up to 2160p, and priority on the GPU queue. Cancelling keeps Premium until the end of the paid period, with a clear end date and an offer to stay.

Pro sits above Premium at 15 USD or 15 EUR a month, or 150 USD or 150 EUR a year. Visitors choose the currency: Europe sees euros by default, a USD / EUR switch sits on the pricing page, in the billing panel and in the upgrade prompts, the choice is remembered, and a subscription keeps the currency it was bought in. The price is the same number in both currencies, there is no conversion, and no tax is added on top. It raises the daily allowance to 700 operations and 120 AI jobs, keeps Premium's file, length and resolution limits, and carries the largest allowance for Motion, the SVG animation editor: 150 exports a day and the MP4, WebM and PNG sequence video exports. Moving up is immediate, with a prorated charge and the same renewal date; moving down to Premium waits until the end of the paid period, so nothing is refunded mid-cycle. The Stripe price id decides the plan on the server, and an unknown price can never grant Pro. The pricing page and the billing panel read every number and price from the API.

## Frontend

Next.js App Router on Vercel, React 19, Tailwind 4, Radix, TanStack Query, Zustand and Zod. Eighteen translated locales served as i18next shards per feature, hreflang alternates in the sitemap, consent-gated first-party analytics with its own admin dashboard (revenue, activation and return, and a Darkroom page with its funnel, feature use, editor performance and export health), a public changelog with an RSS feed, an ads architecture that is off by default and audited in CI, and per-locale translator deployments where a hired translator edits the live page in context and submits a reviewed GitHub pull request.

## Tech

Go, Gin, Postgres, Redis, FFmpeg, ImageMagick, NVENC, Next.js, React, TypeScript, Tailwind, WebGL2, Web Audio, Zod, WebCodecs, mediabunny, anime.js, Stripe, Whisper, Real-ESRGAN, Demucs, Ollama, S3, Firebase Auth, Vercel, nginx, Cloudflare

## Links

- [Visit media-manipulator.com](https://www.media-manipulator.com)
- [Open Darkroom](https://www.media-manipulator.com/tools/content-studio)

## Media

- video: Media Manipulator launch film: a photo converted to sepia, a background removed, GPS and 120 metadata tags stripped, a face blurred, an SVG animated in Motion and a clip edited in Darkroom, ending on the aperture logo and media-manipulator.com — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/media-manipulator/MM_Launch_90.mp4
- image: Media Manipulator home page: the upload panel, quick-tool chips and the Darkroom entry point — https://www.winapps.io/media-manipulator/home.webp
- image: Tool catalogue listing Darkroom first, then image tools such as compress, resize and background removal — https://www.winapps.io/media-manipulator/tools.webp
- image: Darkroom landing page with its quick guide to the multi-track timeline, supported media and export — https://www.winapps.io/media-manipulator/darkroom.webp
- image: AI image restoration tool describing its pre-clean, upscaler and face-enhancement model pipeline — https://www.winapps.io/media-manipulator/restoration.webp
- image: AI document scan tool that turns printed pages and handwriting into a searchable PDF or Word file — https://www.winapps.io/media-manipulator/docscan.webp

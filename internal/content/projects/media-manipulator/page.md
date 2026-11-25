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
`projects/media-manipulator`, `Media Manipulator`, `Go`, `Gin`, `Postgres`, `Redis`, `FFmpeg`, `ImageMagick`, `NVENC`, `Next.js`, `React`, `TypeScript`, `Tailwind`, `WebGL2`, `Web Audio`, `Zod`, `WebCodecs`, `mediabunny`, `anime.js`, `NumPy`, `OpenCV`, `LaMa`, `Stripe`, `Whisper`, `Real-ESRGAN`, `Demucs`, `Ollama`, `Chatterbox`, `Qwen3-TTS`, `FastAPI`, `rubberband`, `S3`, `Firebase Auth`, `Expo`, `React Native`, `Redux Toolkit`, `Firebase Cloud Messaging`, `Crashlytics`, `Vercel`, `nginx`, `Cloudflare`, `Darkroom`, `Content Studio`, `video editor`, `media converter`, `AI restoration`, `document scan`

## Important Links
- https://www.media-manipulator.com
- https://www.media-manipulator.com/tools/content-studio

## Page Content Snapshot

# Media Manipulator

> 95+ browser media tools, a video editor, an SVG animator and a photo editor, on GPU servers I run myself.

- Status: shipped
- Role: Founder, sole engineer
- Dates: May 2025 – Present
- Page: https://www.winapps.io/projects/media-manipulator

Media Manipulator converts, compresses, trims, transcribes, restores and edits images, video and audio in the browser, with the heavy work done on my own GPU servers so large files and AI models are no problem. Darkroom, its multi-track timeline editor, brings colour grading, chroma key, keyframes, transitions, titles and captions to a web page and exports through ffmpeg with preview-to-export parity.

The AI tools (restoration, upscaling, transcription, video summaries, document scanning, face and PII redaction) run on hardware I operate. No user file goes to a third-party AI provider.

## In numbers

| Metric | Value |
| --- | --- |
| Tools | 95+ |
| Languages | 18 |
| Reviewed languages | 4 |
| Transition types | 15 |
| Metadata tags catalogued | 1,800+ |
| Blend modes | 8 |
| Motion export formats | 11 |
| Motion animators | 6 |
| Pro Motion exports a day | 150 |
| Aperture adjustment types | 27 |
| Aperture mask kinds | 6 |
| Aperture golden renders | 126 |
| AI models served | 15+ |
| Voice languages | 23 |
| API routes | ~180 |
| Go tests | 900+ |

## What it is

Media Manipulator started as a converter and became a media workbench. It has 95+ single-purpose tools for images, video and audio, each with its own page, plus Darkroom, a multi-track video editor that runs in a browser tab. Files are processed on servers I run. No account is needed for the standard tools, and guests' files are deleted within 24 hours; signed-in users choose their own retention, from 24 hours to 2 years.

Motion, an SVG animation editor in beta, sits beside Darkroom: it turns an SVG into layers on a keyframe timeline and exports animated SVG, Lottie, React, JavaScript, Flutter or GIF straight from the browser. Aperture, a photo editor in beta, completes the set with adjustment layers, masks and healing, exported at full resolution on the server.

The site is available in 18 languages, 4 reviewed. Every tool's name and one-line description, the Pricing, Account and Voices pages and the site chrome are translated in all 18; every tool page is translated in full in all 18, down to the worked examples and FAQ (the Aperture photo pages are still being finished in a few languages), and English, Ukrainian, Russian and German are reviewed by native speakers. Hebrew, Spanish for Spain and Latin America, French, Italian, Portuguese for Portugal and Brazil, Japanese, Simplified and Traditional Chinese, Korean, Arabic, Hindi and Vietnamese are machine-translated and stay out of search engines (noindex, no hreflang) until a native speaker signs each one off; a tool page whose body would fall back to English is never indexed under a translated address. Two of the languages are right to left, and a key-parity test keeps any translation from silently lagging. It is a CreaTV Ltd. product, and Darkroom is embedded directly in CreaTV Studio.

The About page names me as founder and sole engineer, with CreaTV Ltd. as the operating company, and a contact page offers a short form (spam-limited by a honeypot and a per-session and per-IP rate limit, delivered to the support inbox through Resend with a reference number) next to the support and privacy addresses. The structured data matches what is on the page: a named author on every article and tutorial, a price range that agrees with the pricing page, and sitemap dates taken from the real edit history.

## Darkroom: a timeline editor in the browser

Darkroom composites video, image, audio and LUT tracks in a WebGL2 pipeline: crop, fit, transform, colour grade, 3D LUT, chroma key, opacity and blend, in that order, with a CSS-filter fallback if the GPU context is lost. Audio mixes through a Web Audio graph. Decoders are pooled and Safari's decoder cap is respected.

Editing features: trim, split and ripple delete at frame accuracy; a media bin with server-rendered proxies (one keyframe a second, so preview seeks stay cheap, and a one-click rebuild for clips imported before that), filmstrips and waveform peaks; a Lumetri-style grade with exposure, contrast, saturation, temperature, tint and vibrance; chroma key with an eyedropper; eight blend modes; fifteen transitions; keyframes on position, scale, rotation, opacity, volume, pan and any effect parameter with five easing curves; title clips with text and shape layers; audio ducking and loudness presets; a captions lane with Whisper auto-captions and SRT/VTT export or burn-in.

Captions are word-timed and animated: karaoke fill, a highlighted current word, a pop or a bounce as each word lands, words revealed one by one or typed out letter by letter, a coloured box or a glow behind the spoken word, or a fade per line, in twelve one-click looks and 27 caption fonts with script fallback for Cyrillic, Hebrew, Arabic and CJK. One `.ass` generator, written twice (Go and TypeScript) and held byte-identical by 30 golden fixtures, feeds libass on both sides: ffmpeg burns it into the export and JASSUB (libass compiled to WebAssembly, loaded only when captions show) draws the editor preview, so the preview matches the render to SSIM 0.996+ over the caption.

Paying plans get a per-word caption editor on the timeline: drag the boundary between two words to retime them to the frame, click a word to emphasise it, and restore the transcribed timing without another AI job. Their own caption looks are saved to the account (ten on Premium, unlimited on Pro) under the same compare-and-set, retention-ledger and data-export rules as every other account setting, and the server strips emphasis from the export of a plan that lacks the editor.

Still images can be imported and held on the timeline for any length, Premiere-style, and they export like any other clip. Clips can be sped up or slowed down from 0.1x to 10x with pitch kept, reversed, or frozen on a frame. The timeline has a full keyboard-driven tool set (razor, slip, slide, roll and ripple edits, markers, groups, track lock, solo and hide, and detaching audio from video), and a mixer gives each track a fader, pan, three-band EQ, a compressor and live meters, with a master bus.

The media bin has search, folders, colour labels and rename, a safe delete that lists every clip still using a file, a parallel upload queue, voiceover recorded straight from the microphone onto the timeline, and a Generate voiceover dialog that speaks a typed script in a saved or preset voice (23 languages) on the self-hosted voice engine and drops the clip at the playhead.

Projects are safe to work in: snapshot undo/redo, revision-numbered saves that refuse to clobber, crash-recovery drafts in IndexedDB, and version history with checkpoints. Export is a server-side ffmpeg graph with hardware encoding when available, and its outputs come from one server table of containers and codecs that the dialog, the plan gate and the Medi assistant all read: MP4 with H.264 or H.265, MOV for Final Cut and iMovie, WebM with VP9 and Opus for websites, MKV for archives, or MP3, WAV and Opus audio, plus Pro-only AV1 (GPU-encoded on NVENC, with a CPU fallback behind the encoder limit), ProRes 422 MOV for handing off to Final Cut, Premiere or Resolve (with a size estimate and a server guard that refuses an export over the plan's per-file cap before it renders) and lossless FLAC and 24-bit WAV, with presets for everyone and, for Premium and Pro, an Advanced picker for the container, video and audio codec (AAC, Opus, Vorbis or PCM), audio bitrate and mono (locked options stay visible with the plan that unlocks them), at source resolution or 720p to 2160p, with captions burned in, as SRT/VTT, or embedded as a subtitle track viewers turn on and off; CPU-bound VP9 and ProRes encodes wait their turn behind a server-wide encoder limit and show as Queued, canvas presets for 16:9, 9:16, 1:1 and 4:5, a frame grab as PNG, and an exports panel with progress and re-render. A golden-image harness keeps the preview and the export within a documented tolerance, and solo and hide apply to the export exactly as they do to the preview.

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

## Aperture: a non-destructive photo editor

Aperture is the third editor in the family, in beta, and it is built around a document rather than around pixels. A versioned, additive-only PhotoDoc lists the source, the geometry, the healed areas and an ordered stack of adjustment layers, each with an opacity, a blend mode and an optional mask. It is validated by Zod in the browser and by a Go sanitizer on the server against the same byte-parity fixtures, so the original file is never touched and any step can be reordered, faded or deleted later. Presets sync to your account: signed-in users keep their own looks on the server, capped per plan and validated by the same layer sanitizer, and a one-time prompt moves the presets already saved in the browser.

Three renderers implement one specification. The reference is plain TypeScript in colorMath.ts, with a unit-tested function per adjustment; the WebGL2 preview mirrors it stage for stage in GLSL on a 1920 px copy; the export is rendered from the full-resolution original on the server in NumPy and OpenCV. A golden-image harness renders 126 fixtures through the reference and compares the server's output with tolerances written down per kind of operation (colour operations within a few 8-bit levels, grain exact), so what the preview shows is what the export delivers. Only exposure and white balance work in linear light; the rest are gamma-space stages that clamp to the unit range.

The adjustment set covers exposure, brightness and contrast, a luminance-weighted tone stage, curves (a monotone cubic through up to 16 points per channel), levels, white balance, vibrance, HSL with six ranges plus custom ones picked with an eyedropper, colour balance, black and white with six channel sliders, channel mixer, photo filter, split toning, 3D LUTs from .cube files, posterize, threshold, invert, and a detail group (unsharp mask on luma, a bilateral noise reducer on Lab, clarity, texture, a dark-channel dehaze, vignette and a hash-based grain). A colour-pop look turns every range grey except one, from a single preset.

Layers are limited by masks of six kinds: linear gradient, radial, brush, luminosity range, CIE Lab colour range, and an AI subject selection, folded with add, subtract or intersect and then feathered with a Gaussian. The subject mask and the Heal tool run on my own GPU (rembg and LaMa), so a photo never goes to a third-party service. Geometry is one inverse-mapped resample for crop, straighten, flip and four-corner perspective. Eighteen built-in looks, browser-stored user presets with JSON import and export, and a batch apply that renders one project per photo and zips the results complete the workflow.

Accounts decide what persists and what costs the server, not what the editor can do: projects are 5, 50 or unlimited by plan, export size is 12 megapixels on Free and up to 120 on Premium and Pro, and batches are 3, 10 or 25 photos. Without an account one photo lives in the browser, and only the export, subject mask or heal uploads it for that job. The 24 adjustment pages on the site each open the editor with that panel in front, and a tutorial and a changelog entry document the beta.

- One document, three renderers: TypeScript reference, WebGL2 preview, NumPy/OpenCV export
- 126 golden renders keep the browser and the server within documented tolerances
- Six mask kinds, including a CIE Lab colour range and an AI subject selection on my own GPU
- Export size and batch size follow the plan table, served live from the tier API

## See the result before you download it

Every tool shows its output in the browser before anything is downloaded, with viewers for images, video, GIF and audio. Compare modes include a wipe slider, side-by-side, synced A/B video with an fps readout, and audio A/B. Transcripts and captions open in viewers with click-to-seek and a two-column translation view; there is also a ZIP archive gallery, a PDF viewer, and HLS and DASH stream playback. For AI video restoration, the original can be compared against each model's output.

## Metadata inspector

The inspector reads EXIF, XMP, IPTC and QuickTime metadata through exiftool and ffprobe and explains it instead of dumping raw tags. It is backed by a catalogue of about 1,800 tags in 22 categories, flags what matters for privacy (location, identity, device, history), plots GPS coordinates on a locator, and shows a before/after diff in split or unified view. "Strip everything" covers video and audio as well as images.

## The tool catalogue

The catalogue page has a search box and category tabs; the two dozen most-used tools, each with a page rewritten from the code that runs it (two worked examples with real numbers, a "when not to use this" paragraph and a privacy note specific to that tool, held to a word-count and duplication test in CI), are shown first and the rest sit in a collapsible list per category. Duplicate names for the same tool (MP4 trimmer, video cutter and the like) are thin pages that point to the main one, and search finds a tool by any of its names. Image: format conversion across JPG, PNG, WebP, HEIC, AVIF, SVG and ICO; resize, crop, rotate, compress, sharpen, blur, filters, text overlay, metadata editing, background removal, image-to-PDF and PDF-to-image. Video: conversion between MP4, MOV, WebM, AVI, MKV and HEVC; compress, trim, crop, resize, rotate, flip, reverse, speed change, frame-rate change, stabilisation and colour adjustment; audio extraction and replacement; frame extraction; GIF; HLS packaging with an H.264 ladder and DASH packaging with AV1/VP9, multilingual captions, storyboards and an I-frame playlist. Audio: conversion to MP3, WAV, FLAC, OGG, M4A and Opus; trim, speed, pitch, normalisation, volume, equaliser, reverb, waveform rendering, vocal isolation and removal (the stem comes back in the input's format, so an MP3 stays a 320 kbps MP3 instead of a WAV ten times its size, with FLAC, WAV and a free server-side “Download as MP3” for older WAV results), voice cleanup and background-noise removal. A metadata inspector reads EXIF, GPS, codec, ID3 and PDF tags from any file.

Privacy: EXIF, GPS, IPTC and XMP removal; face blur; text and PII redaction; object removal.

Seven of the tools are recent AI additions: upscale, blur faces, redact text, remove object, clean voice, remove background noise and remove vocals.

## AI on hardware I run

The AI tools run on my own GPU server, so no user file is sent to a third-party model provider. Transcription and captions use Whisper; caption translation uses a TranslateGemma model served by Ollama; video restoration offers six models side by side; image restoration combines upscalers with face-restoration and pre-cleaning models; document scanning turns printed pages into a searchable PDF through Tesseract and transcribes handwriting with a vision model that flags uncertain passages instead of guessing. A VRAM-aware scheduler splits jobs across two GPUs. Every tool that runs a model counts against its own daily AI allowance, kept apart from the everyday one so the two add up: Premium's 300 everyday and 100 AI jobs are 400 jobs a day (3, 10 and 200 AI jobs without an account, on a free account and on Pro), and the nav shows both meters side by side.

Video and audio summaries chain three models into one staged job. Whisper transcribes the speech while ffmpeg picks frames at scene changes, capped by the plan's frame budget (24, 48 or 96); a Qwen3-VL vision model reads them in timestamped batches of four with the transcript lines of that window; then a Qwen text model writes a title, a short summary, chapters with start and end times, topics and the on-screen text, as schema-validated JSON with one repair retry. The result viewer renders it as a document whose chapters seek the original file, with a Markdown copy that matches the server's file byte for byte, and the progress list shows each of the six stages as it runs. A larger deep-synthesis model for Pro sits behind a benchmark gate.

- Restoration: Real-ESRGAN, SwinIR, HAT, BasicVSR++, RVRT, VRT
- Faces and cleanup: GFPGAN, CodeFormer, FBCNN, SCUNet, NAFNet
- Also: RIFE frame interpolation, Demucs vocal isolation, DeepFilterNet denoise, LaMa inpainting, rembg

## Voices: clone, design and change a voice

Three voice tools run on self-hosted models with commercial-use licences. Voice cloning makes a saved voice from 5 to 60 seconds of speech recorded in the browser or uploaded, then reads any text in 23 languages with an emotion control. The voice designer turns a written description (gender, age, accent, tone, pace) into free previews and keeps the one you pick by its seed. The voice changer either shifts pitch and formant instantly with rubberband, or converts a recording to a saved or preset voice with AI, with a synced A/B player to compare.

The models are Chatterbox Multilingual (MIT) for cloning, speech and voice conversion, and Qwen3-TTS (Apache 2.0) for voice design and preset voices. A designed voice is stored as its description, seed and a 10-second rendered reference, and from then on Chatterbox speaks it, so one voice sounds the same everywhere. Both engines live in one small resident FastAPI service on the GPU host that loads models on first use and unloads them when idle, behind a VRAM-aware GPU lease; the page says when the engine is warming up.

Consent and safety are built in rather than bolted on. Creating a voice from a recording requires an explicit confirmation that it is your voice or that you have the speaker's permission, stored with the voice together with the version of that text. Every generated or converted file carries an inaudible PerTh watermark; shifted audio is plain signal processing and is not marked. Recordings never leave the servers I run and are never used to train anything; a voice can be deleted from the account page at any time and is deleted with the account, and the Terms forbid impersonation, fraud and political deepfakes.

- Saved voices by plan: 1 free, 5 Premium, 25 Pro; anonymous visitors can use the shifter and the preset voices
- A failed or refused request gives its daily unit back; design previews are free, three per description an hour
- Results open in the shared result viewer with MP3 and WAV downloads and audio A/B compare

## Medi: an assistant that knows the whole site

Medi is the site's own assistant. On every public page a small button opens a guide that knows every tool, page and editor: ask which tool does a job or where something lives, and it answers in your language with buttons that open the right tool, already set up when the tool has presets. Its knowledge is generated from the site's own registries by a script (every tool page, route, editor parameter, preset and shortcut), searched with BM25 on the API, and never edited by hand, so it stays current as features ship.

The guide is free and runs on Claude Haiku 5.5 through OpenRouter, routed only to Anthropic with data collection denied, behind its own key, rate limits and a 10 USD daily spending cap; at the cap it answers from the help pages alone instead of calling a model. Off-topic questions are turned away before any model call. Signed-in users keep their conversations under the retention period they chose, guests' chats are not stored, and every turn records only content-free usage and cost.

Inside Aperture, Medi is also an editing copilot: say "make everything grey except the red jacket" or "it's backlit, fix the lighting" and it makes the edit, as one step you can undo. The model never writes the photo document; it calls a closed set of nine operations (add or update a layer, grey everything except a colour, mask part of the photo, apply a look, auto tone, crop and straighten), checked against the same JSON Schema on the server before streaming and again in the browser before anything changes, then applied all or nothing through the editor's own code. It sees a small preview frame when you allow it, checks its own result once, and offers anything that costs an AI edit (subject mask, heal) as a button you click.

In Darkroom the same copilot edits video: "make it cinematic", "add a lower-third with my name that slides in", "make the voice clearer and duck the music". It works through eighteen timeline operations (looks and grades, designed titles with entrance and exit animation, transitions, keyframes for Ken Burns moves and fades, speed and freeze frames, audio presets and ducking, splits, trims and markers) that land as one undo step, refuse anything on a locked track, and pass the server's own project sanitizer unchanged. Captions, voiceovers, voice cleanup and export stay buttons you click. The twelve looks it uses are also a Looks row in Darkroom's Effects panel, built only from the existing colour-grade parameters so the preview and the export match.

In Motion, Medi animates: "a lower-third that slides in", "reveal my logo", "pop this title in". It works through eleven operations on the SVG document (ready-made templates, explicit keyframes, the draw-on, motion-path, morph, gradient, mask and blur animators, new shapes and text, SVG import through the editor's own sanitiser, and the composition's length and loop), all landing as one undo step, and it suggests the right export for the destination (Lottie for an app, MP4 for social). The six templates it uses (lower third, logo reveal, title pop, bullet reveal, counter, badge bounce) are also a Templates menu in Motion's top bar, and every one produces a document the server's sanitizer keeps byte for byte.

The guide takes a file too: drop one with the paperclip and it reads the file's type, size, duration and codecs to recommend a tool. Say what you want, or drop a file and say "just run it", and Medi prepares a single-tool job and shows a confirmation card with the settings in your language, the queue length and the allowance it would use. It never starts anything itself: a click is always required. Start runs in place for video compress, resize, rotate, crop and trim, audio extract and remove, and transcription; every other tool opens with Medi's settings prefilled. The settings are validated against the real option schema of about 90 tools, so an unsupported value is refused instead of changed, and the confirmation rate (confirmed over proposed) is tracked.

Edits explain and verify themselves. Every operation carries a one-sentence reason shown on the change card; "review before applying" stages each batch as a before and after card with Apply and Discard, and Show me opens the control an edit moved. After an edit Medi can measure the preview (brightness and clipped shadows and highlights for pictures, integrated loudness and true peak for audio, computed server-side) and adjust for up to three rounds, citing the figure each time. Thumbs up or down on answers, and the share of edits still applied at export, feed the quality numbers.

It is built to be measured, not trusted. Every build has a prompt version: the content hash of the persona, capsules, knowledge docs and operation schemas, stored on each message and usage row, so spend and quality are grouped by version in the admin Medi page (spend against the daily cap, cost per surface and model tier, cache ratio, undo within 60 seconds and thumbs-down with red flags, the allowance-to-checkout funnel, job confirmation). Behaviour is checked by code-graded evals (a few hundred cases across the guide, three copilots, tool routing, tool options, job proposals and the measure-and-adjust loop) that run against the live prompt for cents, and each copilot must clear 80 percent, the guide 90, before it is switched on.

- One streamed endpoint for every surface, with a tool loop for help search and actions validated against a closed list of routes
- In 18 languages, with every guide event in the first-party analytics
- Aperture copilot: one undo step per edit, a plain-words list of what changed, and "Plan my edit" for multi-step looks
- Darkroom copilot and a Looks catalog of twelve one-click grades in the Effects panel
- Motion copilot and a Templates menu of six ready-made animations
- Single-tool jobs: a confirmation card with settings, queue and allowance, started only by a click, for about 90 tools
- Reasons on every edit, review-before-applying, Show me, and measured adjust rounds
- Content-hashed prompt versions, an admin page for spend and quality, and code-graded evals as the go-live gate

## Backend

A Go API (Gin, pgx, Redis) with roughly 180 routes, Postgres with versioned migrations, presigned S3 uploads for large video, a job manager with polling and server-sent events, tiered quotas with per-route rate limits (a job the server refuses or that fails gives its daily unit back), Stripe billing for the Premium tier, user-selectable data retention with export and delete, transactional email through Resend, a public status page fed by a health endpoint, Prometheus metrics, and FFmpeg, ImageMagick, ExifTool, Poppler and friends behind it. It runs as a systemd service on a headless Ubuntu workstation behind nginx and Cloudflare, with 900+ Go tests.

## Accounts, Premium and Pro

Accounts use Firebase Auth: you can sign in with Google, Apple or email, with password reset by email and linking when one address has several sign-in methods. The account screen has tabs for history, projects, billing, privacy and data, and profile; users can export their data or delete their account, share a result through an expiring, revocable link, and choose which emails they get. Guests get a browser-side Recent history. Premium is 8 USD or 8 EUR a month, or 80 USD or 80 EUR a year, through Stripe, and brings larger files, more jobs and exports, output up to 2160p, and priority on the GPU queue. Cancelling keeps Premium until the end of the paid period, with a clear end date and an offer to stay.

Pro sits above Premium at 15 USD or 15 EUR a month, or 150 USD or 150 EUR a year. Visitors choose the currency: Europe sees euros by default, a USD / EUR switch sits on the pricing page, in the billing panel and in the upgrade prompts, the choice is remembered, and a subscription keeps the currency it was bought in. The price is the same number in both currencies, there is no conversion, and no tax is added on top. It raises the daily allowances to 700 everyday jobs and 200 AI jobs, keeps Premium's file, length and resolution limits, and carries the largest allowance for Motion, the SVG animation editor: 150 exports a day and the MP4, WebM and PNG sequence video exports. Moving up is immediate, with a prorated charge and the same renewal date; moving down to Premium waits until the end of the paid period, so nothing is refunded mid-cycle. The Stripe price id decides the plan on the server, and an unknown price can never grant Pro. The pricing page and the billing panel read every number and price from the API.

## The Android app

A native Android app, in closed testing, built with Expo and React Native in TypeScript, with Redux Toolkit and RTK Query on top of the same Go API and the same account as the website: it is signed-in only, so plans, limits, history and retention carry over. Files come in through Android's system pickers or the share sheet, the server-side tools run as tracked jobs, and a push notification through Firebase Cloud Messaging arrives when a long job finishes. The app is available in the same 18 languages, including right-to-left layouts for Arabic and Hebrew. The Aperture photo editor runs on the phone with the same document and shaders ported to Skia: the shared PhotoDoc, reference maths and golden fixtures are copied from the website, every GLSL stage has a Skia shader that is checked against the same goldens in CI, and exports are still rendered on the server; History reopens a past export for editing, a Projects list holds the saved projects, and the sixteen adjustment pages are rows in the Tools tab. Voices come to the phone too: a microphone recorder with a live level meter and a 60-second auto-stop makes a reference recording for cloning (with the same explicit consent record and its version as the website), the designer previews a described voice before keeping it, My voices renames and deletes saved voices, and speech, AI voice conversion and the instant pitch shifter run as tracked jobs, with an Original and Result toggle that keeps the playback position. The Voices tools appear only for the plans and server settings that allow them, and the shifter works even when the voice engine is off.

It has first-party analytics that share the website's service and catalog and add a platform dimension, opt-out and erase controls in Account, and Firebase Crashlytics for crash reports. Code updates ship over the air with EAS Update on a native-fingerprint runtime version, and a signed-in user can delete their account from the app, with a public web page for the same request. The Privacy Policy and Terms have a section for the app in all 18 languages.

## Frontend

Next.js App Router on Vercel, React 19, Tailwind 4, Radix, TanStack Query, Zustand and Zod. Eighteen locales served as i18next shards per feature, a per-locale reviewed flag that decides indexing, hreflang alternates in the sitemap for reviewed locales only, consent-gated first-party analytics with its own admin dashboard (revenue, activation and return, a Darkroom page with its funnel, feature use, editor performance and export health, and a Voices page with creations, generations, failure rate and job duration), a public changelog with an RSS feed, an ads architecture that is off by default and audited in CI, and per-locale translator deployments where a hired translator edits the live page in context and submits a reviewed GitHub pull request.

## Tech

Go, Gin, Postgres, Redis, FFmpeg, ImageMagick, NVENC, Next.js, React, TypeScript, Tailwind, WebGL2, Web Audio, Zod, WebCodecs, mediabunny, anime.js, NumPy, OpenCV, LaMa, Stripe, Whisper, Real-ESRGAN, Demucs, Ollama, Chatterbox, Qwen3-TTS, FastAPI, rubberband, S3, Firebase Auth, Expo, React Native, Redux Toolkit, Firebase Cloud Messaging, Crashlytics, Vercel, nginx, Cloudflare

## Links

- [Visit media-manipulator.com](https://www.media-manipulator.com)
- [Open Darkroom](https://www.media-manipulator.com/tools/content-studio)

## Media

- video: Media Manipulator launch film: a photo converted to sepia, a background removed, GPS and 120 metadata tags stripped, a face blurred, a photo edited in Aperture where Medi turns it black and white except the red, an SVG animated in Motion, and a Darkroom edit with a cross dissolve, a black-and-white freeze frame, an intro Medi made and a two-second dip to black, ending on the aperture logo and media-manipulator.com — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/media-manipulator/MM_Launch_150.mp4
- image: Media Manipulator home page: the upload panel, quick-tool chips and the Darkroom entry point — https://www.winapps.io/media-manipulator/home.webp
- image: Tool catalogue with search and category tabs, listing Darkroom and Motion first, then every image, video, audio, AI and privacy tool — https://www.winapps.io/media-manipulator/tools.webp
- image: Darkroom landing page with its quick guide to the multi-track timeline, supported media and export — https://www.winapps.io/media-manipulator/darkroom.webp
- image: AI image restoration tool describing its pre-clean, upscaler and face-enhancement model pipeline — https://www.winapps.io/media-manipulator/restoration.webp
- image: AI document scan tool that turns printed pages and handwriting into a searchable PDF or Word file — https://www.winapps.io/media-manipulator/docscan.webp

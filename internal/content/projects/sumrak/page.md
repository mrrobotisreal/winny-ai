# Сумрак (Sumrak)

- Route: `/projects/sumrak`
- Canonical URL: `https://www.winapps.io/projects/sumrak`
- Category: `projects`

## Purpose
Generated from the winapps.io content layer (winapps_next/content) by `npm run export:winny`. Do not edit by hand.

## Winny Notes
- Status: active; Creator, sole engineer; Aug 2026 – Present.
- Сумрак is private and will never be in an app store; there is no public repo or download.
- Never state dollar figures or revenue numbers.

## Search Keywords
`projects/sumrak`, `Сумрак`, `Сумрак (Sumrak)`, `Expo`, `React Native`, `TypeScript`, `SQLite`, `Drizzle`, `Kotlin`, `sherpa-onnx`, `Piper`, `Whisper`, `FSRS`, `ElevenLabs`, `Go`, `Postgres`, `Firebase Auth`, `Zod`, `Vitest`, `Sumrak`, `Russian learning`, `language learning`, `ТРКИ`, `TORFL`, `creepypasta`, `Alina`

## Important Links

## Page Content Snapshot

# Сумрак (Sumrak)

> An offline-first Russian-learning app built for one student: me.

- Status: active
- Role: Creator, sole engineer
- Dates: Aug 2026 – Present
- Page: https://www.winapps.io/projects/sumrak

Сумрак (Sumrak) is a private Android app (iOS port in progress) that takes me from A1 toward C1 Russian through graded horror fiction, spaced repetition, blind spoken role-plays, and ТРКИ exam prep. Every part of it is mine: the content pipeline, the on-device speech stack, the sync service, and the twenty-five architecture decisions behind them.

It will never be in a store. It exists so I can speak Russian with my fiancée Alina, who is Ukrainian, and her family.

## In numbers

| Metric | Value |
| --- | --- |
| Content packs | 122 |
| Annotated sentences | 26,867 |
| Game modes | 9 |
| On-device voices | 4 |
| Architecture decision records | 25 |
| TypeScript | ~190k lines |
| Shipped in | 2 months |

## What it is

Сумрак (Sumrak) is a Russian-learning app I built for exactly one user. It runs offline on my phone, keeps all data local, and syncs between my two devices through a small service on my home server. The content is graded fiction in the style of internet horror stories, because I will read a creepypasta every night and I will not read a textbook.

The app owns the whole loop: read a story with narration and word-level highlighting, tap any word or phrase into a word bank, review the bank with spaced repetition and nine game modes, speak through blind role-plays judged on-device, write in a journal that gets AI feedback, and work through the official ТРКИ exam ladder. A progress engine decides what needs work next.

## Content pipeline

Every story, exam and course is a pack: a Zod schema shared by the authoring pipeline, the content repo and the app. The pipeline drafts and annotates text, renders narration with ElevenLabs on a professional clone of my own voice, aligns the audio to the token stream (ElevenLabs timestamps first, then a wav2vec2 forced aligner to re-stamp word timings), stitches a full-story track, and publishes with sha256-pinned manifests. The app verifies every file before import.

User data never holds a foreign key into content. Items reference packs by stable string ids, so a pack can be re-imported or removed without touching a single review record.

- 122 packs across A1–C1: 75 story packs, 17 exam packs, 15 course floors, 10 speaking scenarios
- 504 story entries, 26,867 annotated sentences, 4,461 Opus audio files
- 25 series, each with its own bible

## Reader and word bank

The reader highlights each word as the narrator reaches it. Tapping a word opens a lookup; selecting a phrase opens a phrase card; a sentence can reveal its translation. Anything saved lands in the word bank with the sentence it came from.

A word-bank entry can generate a full profile on demand: gender, animacy, stress, the complete declension or conjugation table, word family, collocations and a usage lesson. Generation goes through OpenRouter with a provider × quality × effort profile and a cost estimate built from real receipts, and it is resumable, so a batch of 400 words can run overnight.

## Spaced repetition and games

Scheduling is FSRS, one card per item and direction, with an offline optimiser that refits the parameters from my own review log. The daily session mixes nine modes by weight: flashcards, multiple choice, cloze (only from sentences I have already read, so nothing is spoiled), sentence builder, listening quiz, pronunciation, a 60-second match sprint, a numbers drill and a typing trainer. Leeches get their own inbox.

## On-device speech

All speech runs on the phone inside one Kotlin Expo module built on sherpa-onnx: four Piper voices for offline text-to-speech, a Zipformer transducer for Russian recognition, a second resident Whisper recogniser for English questions mid-conversation, and a WAV-to-Opus transcoder on the platform media codec with no extra dependency. Pronunciation practice scores each word by alignment and allows unlimited retries.

## Blind spoken role-plays

Scenarios are conversations with no text on screen. A turn state machine with endpointing presets listens, an offline judge scores the answer against authored expectations and slots, and a near-miss gets one time-boxed rescue call to a fast online model that can upgrade the verdict but can never block the turn. Spoken meta-commands («что значит…», «как сказать…», «повтори») work mid-scene. The cast is drawn from PNG layers with mouth shapes driven by loudness tracks computed in the pipeline. Every attempt is transcoded to Opus and archived as an encrypted media bundle.

## ТРКИ exam prep and the Hotel

Exam packs model the official structure: subtests, parts, items, with the texts in the same pack. Mock exams are timed, resumable, scored by the SPbU rule, and graded in layers: offline first, then an AI rubric, then self-assessment. The study path is a hotel: one floor per course pack, a lift between floors, hard gates at 80% to pass and 90% to skip, and an unseen mock exam behind the last door.

## Sync, backup and the API

Backups are zero-knowledge: gzip, then AES-256-GCM with a PBKDF2 key, pushed to two targets with ring-buffer retention. Multi-device sync uses SQLite outbox triggers generated from a table registry (a test regenerates the SQL and diffs it), a hybrid logical clock, and a Go service on my home server that resolves last-writer-wins in Postgres with tombstones and an advisory-locked cursor. The service verifies Firebase ID tokens against Google's certificates without the Admin SDK, has two dependencies, and runs three apps from one binary.

## Progress engine, widget, everything else

Around 470 typed local analytics events feed a dashboard that shows vocabulary by level, grammar coverage, a thirty-day forecast, a year heatmap and an AI CEFR assessment, loading section by section on a second SQLite connection so the screen opens on the next frame. A Glance home-screen widget renders a JSON snapshot the app writes, never touching the database. Themed ambient soundtracks are chosen per story. The settings tour is long because every behaviour is mine to tune.

## How it was built

Two months, 463 commits, 25 architecture decision records, a ubiquitous-language glossary per repo, 284 test files, and cross-language contract tests (Kotlin constants against TypeScript, trigger SQL against the registry, the sync allowlist against the app). Content work runs through its own ticket system. An iOS port is in progress, and the same codebase has been forked into a Ukrainian app and an English-for-Russian-speakers app.

## Tech

Expo, React Native, TypeScript, SQLite, Drizzle, Kotlin, sherpa-onnx, Piper, Whisper, FSRS, ElevenLabs, Go, Postgres, Firebase Auth, Zod, Vitest

## Media

- image: Today screen with goal rings, a finished quest and 95 cards due — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_001_540.jpg
- video: Reading a story with narration, word highlighting and lookups — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Reading_540.mp4
- video: A review session moving through cloze, sentence builder and other game modes — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Review_540_trimmed.mp4
- image: Word-bank list filtered by level, with gender and part-of-speech tags — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_015_540.jpg
- image: Hotel path: the lift door on floor 0, the lobby, under the ТРКИ-А1 floor counter — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_002_540.jpg
- image: Library shelves of A1 horror stories with level and progress tags — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_003_540.jpg
- image: The Three-Body series list in the library, filtered to science fiction — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_008_540.jpg
- image: Story chapter with a full-story listen button and its parts — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_009_540.jpg
- image: Reader page with Russian text and a tapped word highlighted — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_010_540.jpg
- image: Reader playing narration, with the current word highlighted in the text — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_011_540.jpg
- image: Story pack with its chapter list and the mini-player — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_012_540.jpg
- image: ТРКИ-А1 and ТРКИ-А2 shelves with readiness bars per subtest — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_013_540.jpg
- image: ТРКИ-А1 lexicon and grammar topic units — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_014_540.jpg
- image: Word-bank entry for туман with review directions and history — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_016_540.jpg
- image: Generating the full grammatical profile for туман — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_018_540.jpg
- image: Generated declension table and word family for туман — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_020_540.jpg
- image: Generated collocations lesson with example sentences for туман — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_027_540.jpg
- image: Scenario hub cards: at the bank, at the doctor, the oral exam — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_052_540.jpg
- image: More scenarios: gas station, night shift, on air, at the police station — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_053_540.jpg
- image: Prompt to install the on-device assist model for English questions — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_054_540.jpg
- image: Progress dashboard: vocabulary by level and core vocabulary coverage — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_030_540.jpg
- image: Progress dashboard: 30-day review forecast and grammar coverage — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_031_540.jpg
- image: Progress dashboard: pronunciation trend, checkpoints, year heatmap and AI assessment — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_033_540.jpg
- image: Home-screen widget with the daily goal ring and cards due — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_Widget_540.jpg
- image: Settings: appearance, reading and scenario options — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_035_540.jpg
- image: Settings: background music, daily goal and streak freezes — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_039_540.jpg
- image: Settings: four on-device Piper voices for Russian speech — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_041_540.jpg
- image: Settings: on-device Russian speech recognition and the Whisper assist model — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_042_540.jpg
- image: Settings: FSRS desired retention and new cards per day — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_044_540.jpg
- image: Settings: content sync from the private content repository — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_045_540.jpg
- image: Settings: encrypted backups to two targets — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_046_540.jpg
- image: Settings: home-server sync target and connection test — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_047_540.jpg
- image: Games list: daily session, review, cloze, sentence builder, listening quiz — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_050_540.jpg
- image: More games: typing trainer, numbers, the 60-second match sprint, dialogue, scenarios — https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/sumrak/Screenshot_051_540.jpg

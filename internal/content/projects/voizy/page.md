# Voizy

- Route: `/projects/voizy`
- Canonical URL: `https://www.winapps.io/projects/voizy`
- Category: `projects`

## Purpose
Generated from the winapps.io content layer (winapps_next/content) by `npm run export:winny`. Do not edit by hand.

## Winny Notes
- Status: paused; Creator, sole engineer; Mar 2025 – Apr 2025.
- Never state dollar figures or revenue numbers.

## Search Keywords
`projects/voizy`, `Voizy`, `Kotlin`, `Jetpack Compose`, `Dagger/Hilt`, `Go`, `Java`, `Jetty`, `MySQL`, `Mockito`, `Python`, `scikit-learn`, `Google Maps API`, `Firebase`, `Epidemic Sound`

## Important Links
- https://github.com/mrrobotisreal/VoizyAndroid
- https://www.winapps.io/projects/voizy/documentation
- https://www.winapps.io/projects/voizy/privacy-policy
- https://www.winapps.io/projects/voizy/support

## Page Content Snapshot

# Voizy

> An Android social app in Kotlin with a Java backend and a Python recommendation service.

- Status: paused
- Role: Creator, sole engineer
- Dates: Mar 2025 – Apr 2025
- Page: https://www.winapps.io/projects/voizy

Voizy is a social app for Android with no ads and no tracking, written in Kotlin and Jetpack Compose. People build a profile with photos, a cover and a profile song, find and add friends, post to the feed or a friend's page with images, locations, hashtags, polls and tags, react and comment, repost, join groups, and tune their notifications.

Behind it are a Go server, a Java server on Jetty, and a Python service that ranks the recommended-posts feed.

## Android app

The client is Kotlin and Jetpack Compose. It covers sign-up and profile management, the social feed with reactions and comments, friend search and requests, image and audio uploads with compression, location posts with Google Maps, groups, and configurable notifications.

JWTs and API keys are kept in Android's secure storage, sensitive local data is encrypted, posts, images and audio are cached, data loads in pages, and dependency injection runs on Dagger/Hilt.

## Backend

The API exists in two versions on MySQL: a Go server, and a Java server on Eclipse Jetty servlets with an MVC structure and Mockito tests. Both handle JWT and API key auth, complex SQL joins over indexed tables, and caching for posts, images and audio.

The recommended-posts feed has a ranking algorithm written in Go, plus a separate Python service built with scikit-learn, pandas and NumPy.

## Integrations

Google Maps provides locations, Epidemic Sound provides profile music, and Firebase delivers push notifications.

## Tech

Kotlin, Jetpack Compose, Dagger/Hilt, Go, Java, Jetty, MySQL, Mockito, Python, scikit-learn, Google Maps API, Firebase, Epidemic Sound

## Links

- [Android source on GitHub](https://github.com/mrrobotisreal/VoizyAndroid)
- [Documentation](https://www.winapps.io/projects/voizy/documentation)
- [Privacy policy](https://www.winapps.io/projects/voizy/privacy-policy)
- [Support](https://www.winapps.io/projects/voizy/support)

## Media

- image: Voizy: Login screen — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_033.jpg
- image: Voizy: Create account screen — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_034.jpg
- image: Voizy: Profile feed with a post containing a photo — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_001.jpg
- image: Voizy: Profile feed with posts from friends — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_002.jpg
- image: Voizy: Comments sheet open on a post — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_003.jpg
- image: Voizy: Writing a comment with the keyboard open — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_004.jpg
- image: Voizy: Post with an image carousel — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_005.jpg
- image: Voizy: Reaction picker on a post — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_006.jpg
- image: Voizy: Profile photos tab while images upload — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_007.jpg
- image: Voizy: Choosing a photo to set as profile or cover picture — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_008.jpg
- image: Voizy: Full-screen photo with set-as-profile and set-as-cover actions — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_009.jpg
- image: Voizy: Cover picture updated confirmation — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_010.jpg
- image: Voizy: Profile photos grid after updating the cover picture — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_011.jpg
- image: Voizy: Full-screen view of a profile photo — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_012.jpg
- image: Voizy: Profile picture updated confirmation — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_013.jpg
- image: Voizy: Add profile images sheet, empty — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_014.jpg
- image: Voizy: Taking a new photo with the camera — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_015.jpg
- image: Voizy: Reviewing a captured photo before keeping it — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_016.jpg
- image: Voizy: Add profile images sheet with a captured photo — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_017.jpg
- image: Voizy: Choosing photos from the device library — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_018.jpg
- image: Voizy: Add profile images sheet with a selected photo — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_019.jpg
- image: Voizy: Uploading profile images — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_020.jpg
- image: Voizy: Profile photos grid after a successful upload — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_021.jpg
- image: Voizy: Profile info tab — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_022.jpg
- image: Voizy: Friends list with search — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_023.jpg
- image: Voizy: Searching friends by name — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_024.jpg
- image: Voizy: Create post screen, empty — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_025.jpg
- image: Voizy: Searching for a location to add to a post — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_026.jpg
- image: Voizy: Picking a location on the map — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_027.jpg
- image: Voizy: Create post with a location added — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_028.jpg
- image: Voizy: Selecting images to attach to a post — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_029.jpg
- image: Voizy: Writing post text with the keyboard open — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_030.jpg
- image: Voizy: Create post with an attached image — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_031.jpg
- image: Voizy: Feed showing the new post after creation — https://pub-2c90fb271b9b4f08bffffbf1179f9627.r2.dev/voizy_032.jpg

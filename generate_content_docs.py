#!/usr/bin/env python3

from __future__ import annotations

import html
import re
import urllib.error
import urllib.request
from pathlib import Path
from typing import Iterable

ROOT = Path(__file__).resolve().parent
APP_ROOT = ROOT.parent.parent / "Frontend" / "winapps_ui" / "winapps_next" / "app"
OUTPUT_ROOT = ROOT / "internal" / "content"
BASE_URL = "https://winapps.io"
USER_AGENT = "WinnyContentGenerator/1.0 (+https://winapps.io)"

DOWNLOAD_AND_REDIRECT_INFO = {
    "/download/creatv-mobile/android": {
        "title": "Download CreaTV Mobile for Android",
        "summary": [
            "This route immediately downloads the current CreaTV Android development APK.",
            "The saved filename is `CreaTV-Android-development-build.apk`.",
            "After the download starts, the page shows a short countdown and redirects back to the homepage.",
        ],
        "links": [
            "https://pub-5e3f5f69f6bd4f2fb6bc741e03f34851.r2.dev/build-1770091566601.apk",
        ],
        "notes": [
            "Treat this as a utility download route, not a long-form product overview page.",
            "If a visitor asks how to get the Android build, mention that this page triggers the APK download automatically.",
        ],
    },
    "/download/mp3-drive-player/android": {
        "title": "Download MP3 Drive Player for Android",
        "summary": [
            "This route immediately downloads the Android APK for MP3 Drive Player.",
            "The download target is hosted on Cloudflare R2.",
            "The page then redirects back to the homepage after a brief branded countdown.",
        ],
        "links": [
            "https://pub-c0247ba91a4a415a9ff6d54583d7667c.r2.dev/MP3-Drive-Player-Android.apk",
        ],
        "notes": [
            "Use this page when someone asks where to download the Android test build for MP3 Drive Player.",
        ],
    },
    "/download/mp3-drive-player/ios": {
        "title": "Download MP3 Drive Player for iOS",
        "summary": [
            "This route immediately downloads the iOS IPA for MP3 Drive Player.",
            "The build is hosted on Cloudflare R2.",
            "The page uses a short branded countdown and then redirects back to the homepage.",
        ],
        "links": [
            "https://pub-c0247ba91a4a415a9ff6d54583d7667c.r2.dev/MP3-Drive-Player-iOS.ipa",
        ],
        "notes": [
            "Use this route when someone asks how to access the iOS test build for MP3 Drive Player.",
        ],
    },
    "/redirects/photos/albums/a-new-circle-of-love-she-said-yes": {
        "title": "Photo Album Redirect: A New Circle Of Love (She Said Yes)",
        "summary": [
            "Branded redirect page for a Google Photos album.",
            "The page shows a short loading / countdown experience before sending the visitor to the album.",
        ],
        "links": ["https://photos.app.goo.gl/AgsnBxncc1BPLdUA7"],
        "notes": [
            "This is a redirect helper page, not a full photo gallery hosted directly on WinApps.io.",
        ],
    },
    "/redirects/photos/albums/austria": {
        "title": "Photo Album Redirect: Austria",
        "summary": [
            "Branded redirect page for the Austria Google Photos album.",
        ],
        "links": ["https://photos.app.goo.gl/cTAzjrJs3TSf1ZAA6"],
        "notes": ["This route exists to redirect visitors to the external album."],
    },
    "/redirects/photos/albums/austria-2": {
        "title": "Photo Album Redirect: Austria 2",
        "summary": [
            "Branded redirect page for a second Austria Google Photos album.",
        ],
        "links": ["https://photos.app.goo.gl/MeYiUEbMo3EXafMd8"],
        "notes": ["This route exists to redirect visitors to the external album."],
    },
    "/redirects/photos/albums/frankfurt": {
        "title": "Photo Album Redirect: Frankfurt",
        "summary": [
            "Branded redirect page for the Frankfurt Google Photos album.",
        ],
        "links": ["https://photos.app.goo.gl/MD1i8LYGrYrzfYmY9"],
        "notes": ["This route exists to redirect visitors to the external album."],
    },
    "/redirects/photos/albums/gabe-christmas-concert-2025": {
        "title": "Photo Album Redirect: Gabe Christmas Concert 2025",
        "summary": [
            "Branded redirect page for the Gabe Christmas Concert 2025 Google Photos album.",
        ],
        "links": ["https://photos.app.goo.gl/SGNHkBcT9WdCuTBx7"],
        "notes": ["This route exists to redirect visitors to the external album."],
    },
    "/redirects/photos/albums/koala-ventures": {
        "title": "Photo Album Redirect: Koala Ventures",
        "summary": [
            "Branded redirect page for the Koala Ventures Google Photos album.",
        ],
        "links": ["https://photos.app.goo.gl/5gq1JMJ9eQXm56U56"],
        "notes": ["This route exists to redirect visitors to the external album."],
    },
    "/redirects/photos/albums/new-york": {
        "title": "Photo Album Redirect: New York",
        "summary": [
            "Branded redirect page for the New York Google Photos album.",
        ],
        "links": ["https://photos.app.goo.gl/1CHWjNXHvTGEbXoA9"],
        "notes": ["This route exists to redirect visitors to the external album."],
    },
    "/redirects/photos/albums/paris": {
        "title": "Photo Album Redirect: Paris",
        "summary": [
            "Branded redirect page for the Paris Google Photos album.",
        ],
        "links": ["https://photos.app.goo.gl/SByo7CssdN7bLDQu5"],
        "notes": ["This route exists to redirect visitors to the external album."],
    },
    "/redirects/photos/albums/portugal": {
        "title": "Photo Album Redirect: Portugal",
        "summary": [
            "Branded redirect page for the Portugal Google Photos album.",
        ],
        "links": ["https://photos.app.goo.gl/zWBkiY3iwobVVc4Q6"],
        "notes": ["This route exists to redirect visitors to the external album."],
    },
    "/redirects/photos/albums/portugal-photoshoot": {
        "title": "Photo Album Redirect: Portugal Photoshoot",
        "summary": [
            "Branded redirect page for the Portugal Photoshoot Google Photos album.",
        ],
        "links": ["https://photos.app.goo.gl/rAUy4cgNwqYEuzvZ9"],
        "notes": ["This route exists to redirect visitors to the external album."],
    },
    "/redirects/photos/albums/rome-stria": {
        "title": "Photo Album Redirect: Rome Stria",
        "summary": [
            "Branded redirect page for the Rome Stria Google Photos album.",
        ],
        "links": ["https://photos.app.goo.gl/dJfwHuEoAvnpkGao9"],
        "notes": ["This route exists to redirect visitors to the external album."],
    },
}


def route_from_page(page_path: Path) -> str:
    relative = page_path.relative_to(APP_ROOT)
    parts = relative.parts[:-1]
    if not parts:
        return "/"
    return "/" + "/".join(parts)


def output_path_for_route(route: str) -> Path:
    if route == "/":
        return OUTPUT_ROOT / "page.md"
    return OUTPUT_ROOT / route.lstrip("/") / "page.md"


def title_from_html(document: str, fallback: str) -> str:
    main_content = main_from_html(document)
    h1_match = re.search(r"(?is)<h1[^>]*>(.*?)</h1>", main_content)
    if h1_match:
        heading = visible_text_from_html(h1_match.group(1))
        if heading:
            return heading
    match = re.search(r"(?is)<title[^>]*>(.*?)</title>", document)
    if match:
        return normalize_whitespace(html.unescape(match.group(1)))
    return fallback


def main_from_html(document: str) -> str:
    match = re.search(r"(?is)<main[^>]*>(.*?)</main>", document)
    if match:
        return match.group(1)
    return document


def visible_text_from_html(document: str) -> str:
    document = main_from_html(document)
    document = re.sub(r"(?is)<head[^>]*>.*?</head>", " ", document)
    document = re.sub(r"(?is)<script[^>]*>.*?</script>", " ", document)
    document = re.sub(r"(?is)<style[^>]*>.*?</style>", " ", document)
    document = re.sub(r"(?is)<noscript[^>]*>.*?</noscript>", " ", document)
    document = re.sub(r"(?is)<svg[^>]*>.*?</svg>", " ", document)
    document = re.sub(r"(?s)<!--.*?-->", " ", document)
    document = re.sub(r"(?s)<[^>]+>", " ", document)
    return normalize_whitespace(html.unescape(document))


def links_from_html(document: str) -> list[str]:
    document = main_from_html(document)
    matches = re.findall(r'''href=["']([^"']+)["']''', document, flags=re.IGNORECASE)
    seen: set[str] = set()
    links: list[str] = []
    for link in matches:
        link = link.strip()
        if not link or link.startswith("#") or link.startswith("mailto:") or link.startswith("tel:"):
            continue
        if any(
            blocked in link
            for blocked in [
                "/_next/",
                "favicon",
                ".woff",
                ".woff2",
                ".css",
                ".js",
            ]
        ):
            continue
        if link.startswith("/"):
            link = BASE_URL.rstrip("/") + link
        if link in seen:
            continue
        seen.add(link)
        links.append(link)
    return links


def emails_from_text(value: str) -> list[str]:
    return sorted(set(re.findall(r"[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}", value)))


def normalize_whitespace(value: str) -> str:
    return re.sub(r"\s+", " ", value).strip()


def truncate(value: str, limit: int) -> str:
    value = value.strip()
    if len(value) <= limit:
        return value
    return value[: limit - 3].rstrip() + "..."


def fallback_text_from_source(source_text: str) -> str:
    strings_found = re.findall(r'(?s)"([^"\n]{8,})"|`([^`]{8,})`', source_text)
    flattened: list[str] = []
    for first, second in strings_found:
        value = first or second
        value = normalize_whitespace(value)
        if value and "use client" not in value and not value.startswith("http://localhost"):
            flattened.append(value)
    deduped: list[str] = []
    seen: set[str] = set()
    for value in flattened:
        if value in seen:
            continue
        seen.add(value)
        deduped.append(value)
    return truncate("\n".join(deduped[:80]), 5000)


def fetch_live_page(route: str) -> tuple[str, str, list[str]]:
    url = BASE_URL if route == "/" else BASE_URL.rstrip("/") + route
    request = urllib.request.Request(
        url,
        headers={
            "User-Agent": USER_AGENT,
            "Accept": "text/html,application/xhtml+xml",
        },
    )
    with urllib.request.urlopen(request, timeout=15) as response:
        html_text = response.read().decode("utf-8", errors="replace")
    title = title_from_html(html_text, route or "/")
    text = visible_text_from_html(html_text)
    links = links_from_html(html_text)
    return title, text, links


def route_category(route: str) -> str:
    if route == "/":
        return "home"
    if route.startswith("/about"):
        return "about"
    if route.startswith("/blog"):
        return "blog"
    if route.startswith("/contact"):
        return "contact"
    if route.startswith("/portfolio"):
        return "portfolio"
    if route.startswith("/projects"):
        return "projects"
    if route.startswith("/products"):
        return "products"
    if route.startswith("/download"):
        return "downloads"
    if route.startswith("/redirects"):
        return "redirects"
    return "other"


def generic_notes(route: str) -> list[str]:
    category = route_category(route)
    notes: list[str] = []
    if category == "home":
        notes.extend(
            [
                "Use this page for broad questions about what WinApps.io contains and where visitors should navigate next.",
                "This is a good entry point when someone asks for a quick overview of Mitch, his skills, projects, and portfolio.",
            ]
        )
    elif category == "about":
        notes.extend(
            [
                "Use this page when someone asks about Mitchell Wintrow's background, work history, education, or technical skill set.",
            ]
        )
    elif category == "blog":
        notes.extend(
            [
                "This page is a curated external article hub, so clarify that the articles themselves live off-site when appropriate.",
            ]
        )
    elif category == "contact":
        notes.extend(
            [
                "Use this page for direct contact methods, social links, and where visitors should reach Mitch.",
            ]
        )
    elif category == "portfolio":
        notes.extend(
            [
                "Portfolio pages describe Mitchell's professional work history and project contributions for employers or clients.",
            ]
        )
    elif category == "projects":
        notes.extend(
            [
                "Project pages usually contain the deepest technical and product details for WinApps.io offerings.",
                "When the visitor asks about tech stack, architecture, features, or live demos, these pages are often the best sources.",
            ]
        )
    elif category == "products":
        notes.extend(
            [
                "Product pages are customer-facing and may overlap with project pages while emphasizing support, policy, and end-user information.",
            ]
        )
    elif category == "downloads":
        notes.extend(
            [
                "These pages are operational download launchers and not long-form documentation pages.",
            ]
        )
    elif category == "redirects":
        notes.extend(
            [
                "These routes are redirect helpers to Google Photos albums; they are not fully hosted album pages.",
            ]
        )

    if route == "/about/mitchell-wintrow":
        notes.append("This page mirrors the main `/about` content very closely.")
    if route == "/products/journey-app":
        notes.append("This page substantially overlaps with `/projects/journey-app` and can be treated as the customer-facing counterpart.")
    if route == "/projects/voizy/documentation":
        notes.append("This page primarily points visitors to a Notion-hosted documentation experience.")
    if route == "/projects/text-formatter/privacy-policy":
        notes.append("This policy page is especially useful for privacy, permissions, and extension data-collection questions.")
    if route == "/products/journey-app/contact/form":
        notes.append("Use this page when someone asks about Journey App support categories, file upload limits, or where form submissions are sent.")
    return notes


def search_keywords(route: str, title: str, text: str) -> list[str]:
    values = re.findall(r"[A-Za-z0-9][A-Za-z0-9.+@/-]{1,}", f"{route} {title} {text}")
    seen: set[str] = set()
    keywords: list[str] = []
    for value in values:
        cleaned = value.strip(".,:;()[]{}")
        lowered = cleaned.lower()
        if len(cleaned) < 3:
            continue
        if lowered in {
            "the",
            "and",
            "for",
            "with",
            "this",
            "that",
            "from",
            "page",
            "site",
            "home",
            "https",
            "http",
            "www",
        }:
            continue
        if lowered in seen:
            continue
        seen.add(lowered)
        keywords.append(cleaned)
        if len(keywords) >= 20:
            break
    return keywords


def build_markdown(
    *,
    route: str,
    source_page: str,
    title: str,
    canonical_url: str,
    category: str,
    text_snapshot: str,
    links: Iterable[str],
    emails: Iterable[str],
    notes: Iterable[str],
) -> str:
    lines: list[str] = [
        f"# {title}",
        "",
        f"- Route: `{route}`",
        f"- Source page: `{source_page}`",
        f"- Canonical URL: `{canonical_url}`",
        f"- Category: `{category}`",
        "",
        "## Purpose",
        "Backend content mirror for Winny so the assistant can answer questions about this page quickly and accurately.",
        "",
    ]

    note_list = list(notes)
    if note_list:
        lines.append("## Winny Notes")
        lines.extend(f"- {note}" for note in note_list)
        lines.append("")

    keyword_list = search_keywords(route, title, text_snapshot)
    if keyword_list:
        lines.append("## Search Keywords")
        lines.append(", ".join(f"`{keyword}`" for keyword in keyword_list))
        lines.append("")

    email_list = list(emails)
    if email_list:
        lines.append("## Important Emails")
        lines.extend(f"- `{email_address}`" for email_address in email_list)
        lines.append("")

    link_list = list(links)
    if link_list:
        lines.append("## Important Links")
        lines.extend(f"- {link}" for link in link_list[:20])
        lines.append("")

    lines.append("## Page Content Snapshot")
    lines.append("")
    lines.append(text_snapshot or "No readable page snapshot was available.")
    lines.append("")
    return "\n".join(lines).strip() + "\n"


def clean_generated_docs() -> None:
    if not OUTPUT_ROOT.exists():
        return
    for path in OUTPUT_ROOT.rglob("*"):
        if path.is_file() and path.suffix in {".md", ".txt"}:
            path.unlink()
    for path in sorted(OUTPUT_ROOT.rglob("*"), reverse=True):
        if path.is_dir() and path != OUTPUT_ROOT and not any(path.iterdir()):
            path.rmdir()


def write_sitemaps(entries: list[dict[str, str]]) -> None:
    grouped: dict[str, list[dict[str, str]]] = {}
    for entry in entries:
        grouped.setdefault(entry["category"], []).append(entry)

    sitemap_md_lines = [
        "# Winny Content Sitemap",
        "",
        "This sitemap maps website routes to the mirrored markdown docs under `internal/content/**`.",
        "",
    ]
    for category in sorted(grouped):
        sitemap_md_lines.append(f"## {category.title()}")
        for entry in grouped[category]:
            sitemap_md_lines.append(
                f"- `{entry['route']}` -> `{entry['doc_path']}` ({entry['title']})"
            )
        sitemap_md_lines.append("")

    sitemap_txt_lines = [
        f"{entry['route']} | {entry['title']} | {entry['doc_path']}"
        for entry in entries
    ]

    (OUTPUT_ROOT / "sitemap.md").write_text(
        "\n".join(sitemap_md_lines).strip() + "\n", encoding="utf-8"
    )
    (OUTPUT_ROOT / "sitemap.txt").write_text(
        "\n".join(sitemap_txt_lines).strip() + "\n", encoding="utf-8"
    )


def main() -> None:
    if not APP_ROOT.exists():
        raise SystemExit(f"App root not found: {APP_ROOT}")

    OUTPUT_ROOT.mkdir(parents=True, exist_ok=True)
    clean_generated_docs()

    entries: list[dict[str, str]] = []
    page_files = sorted(APP_ROOT.rglob("page.tsx"))

    for page_file in page_files:
        route = route_from_page(page_file)
        source_page = str(page_file.relative_to(APP_ROOT.parent).as_posix())
        canonical_url = BASE_URL if route == "/" else BASE_URL.rstrip("/") + route
        category = route_category(route)

        title = route or "/"
        text_snapshot = ""
        links: list[str] = []
        notes = generic_notes(route)

        if route in DOWNLOAD_AND_REDIRECT_INFO:
            info = DOWNLOAD_AND_REDIRECT_INFO[route]
            title = info["title"]
            text_snapshot = "\n".join(info["summary"])
            links = list(info.get("links", []))
            notes = [*notes, *info.get("notes", [])]
        else:
            try:
                title, live_text, live_links = fetch_live_page(route)
                text_snapshot = truncate(live_text, 5000)
                links = live_links
            except (urllib.error.URLError, TimeoutError, ValueError):
                source_text = page_file.read_text(encoding="utf-8")
                title = route.strip("/") or "Home"
                text_snapshot = fallback_text_from_source(source_text)

        emails = emails_from_text(text_snapshot + "\n" + "\n".join(links))
        markdown = build_markdown(
            route=route,
            source_page=source_page,
            title=title,
            canonical_url=canonical_url,
            category=category,
            text_snapshot=text_snapshot,
            links=links,
            emails=emails,
            notes=notes,
        )

        output_path = output_path_for_route(route)
        output_path.parent.mkdir(parents=True, exist_ok=True)
        output_path.write_text(markdown, encoding="utf-8")

        entries.append(
            {
                "route": route,
                "title": title,
                "category": category,
                "doc_path": str(output_path.relative_to(ROOT).as_posix()),
            }
        )

    entries.sort(key=lambda item: item["route"])
    write_sitemaps(entries)


if __name__ == "__main__":
    main()

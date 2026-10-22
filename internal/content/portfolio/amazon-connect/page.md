# Software Development Engineer (Front End), Amazon Connect Customer Profiles, Amazon Web Services

- Route: `/portfolio/amazon-connect`
- Canonical URL: `https://www.winapps.io/portfolio/amazon-connect`
- Category: `portfolio`

## Purpose
Generated from the winapps.io content layer (winapps_next/content) by `npm run export:winny`. Do not edit by hand.

## Winny Notes
- Status: Software Development Engineer (Front End), Amazon Connect Customer Profiles at Amazon Web Services, Aug 2022 – Nov 2024 (Full-time).
- This role has ended (Nov 2024). Never describe it as current.

## Search Keywords
`portfolio/amazon-connect`, `Amazon Web Services`, `Amazon Connect`, `Software Development Engineer (Front End), Amazon Connect Customer Profiles`, `React`, `TypeScript`, `React Query`, `Node.js`, `Java 11`, `Python`, `AWS`, `CloudWatch`, `Customer Profiles`, `Customer Segments`, `contact record template`

## Important Links
- https://docs.aws.amazon.com/connect/latest/adminguide/inferred-profiles.html
- https://docs.aws.amazon.com/connect/latest/adminguide/ctr-contact-record-template.html

## Page Content Snapshot

# Software Development Engineer (Front End), Amazon Connect Customer Profiles, Amazon Web Services

> Front-end engineer on Amazon Connect Customer Profiles, building AWS Console and Connect admin website features for contact-center customers.

- Dates: Aug 2022 – Nov 2024
- Employment: Full-time
- Page: https://www.winapps.io/portfolio/amazon-connect

## Highlights

- Built the CTR-AutoAssociateOnly contact record template, which looks up a caller's unified profile by phone number without creating duplicate profiles.
- Built the AWS Console settings for the new template and rewrote its public documentation, adding a section for the auto-associate flow.
- Built the Customer Segments list, segment details and CSV export in the Connect admin website.
- Raised test coverage from 64% to 91% and made 90% coverage a required gate for code reviews.
- Cut false on-call alarms with CloudWatch anomaly detection on 4xx errors, with 403s split into their own dashboard and alarm.
- Wrote a demo-environment script that provisions a domain with 1,000,000 mock profiles, and a Node.js script that inserts new translation strings into every locale file.

## Role

I was a front-end engineer on the Customer Profiles team. I worked in React and TypeScript, with Node.js, Java 11 and Python where a feature needed it.

- Took feature specifications through to maintainable, extensible code.
- Improved existing features independently, with minimal guidance.
- Reviewed code and gave feedback to other engineers.
- Worked with product managers to turn customer requirements into solutions.
- Handled operational support and drove root causes to resolution.
- Improved team processes and tools, ran knowledge-sharing sessions, and trained new front-end engineers.

## Contact record templates

Both existing contact trace record (CTR) templates created a new customer profile automatically, which led to duplicate profiles and inconsistent data. Customers had raised it many times. I built a new template, `CTR-AutoAssociateOnly`. It still looks up the caller's unified profile by incoming phone number in real time, but it never creates a new one.

I then updated the public documentation. I added a section on associating contact records with the `_phone` key and diagrams of the CTR flows. I also rewrote the existing template docs, which customers had called confusing.

Finally I built the front end. The profile creation and auto-association section appears on the Create domain page, and on the Domain details page for viewing and changing the applied template. An info help panel explains each template and links to the documentation.

## Customer Segments

In the Connect admin website I designed and built the Customer Segments landing page. It shows a searchable, filterable table of a customer's segments, or an empty state with a Create Segment button. Segments come from a paginated API and are cached with React Query to keep the table in sync and reduce API calls.

The segment details page shows the profile count, name, ID, description, creation date, and the audience definition (audience groups, filter groups, and include/exclude operators). From there customers can delete, duplicate, or export the segment.

Export to CSV is available from both pages and covers up to 350,000 profiles. Larger exports are left to the AWS Console, because a file that size can run to several gigabytes. The console also gives IT admins better controls than the admin website offers contact-center admins.

## Operations and tooling

Setting up a realistic Customer Profiles domain for testing or demos took engineers and product managers a long time. I wrote a script with Bash, PowerShell and Cmd entry points that installs the right Python if needed. It then provisions a domain, enables every feature, and loads 1,000,000 generated profiles and products. The team used it for development and for customer demos.

Our 4xx alarms fired on any spike, so on-call engineers kept chasing customers' permission errors and harmless 404s. I added CloudWatch anomaly detection so only unusual spikes alarm. I also moved 403s to their own dashboard and alarm, to track customers struggling with permissions setup.

Code reviews had no coverage requirement, and simple bugs were reaching production. I made 90% coverage a required gate and raised the codebase from 64% to 91%.

Adding a new UI string meant copying it into every locale file in exactly the right place by hand. I wrote a Node.js script that takes new strings from the English file and inserts them into every other locale. The whole front-end team adopted it.

## Tech

React, TypeScript, React Query, Node.js, Java 11, Python, AWS, CloudWatch

## Links

- [Inferred profiles (AWS docs)](https://docs.aws.amazon.com/connect/latest/adminguide/inferred-profiles.html)
- [Contact record templates (AWS docs)](https://docs.aws.amazon.com/connect/latest/adminguide/ctr-contact-record-template.html)

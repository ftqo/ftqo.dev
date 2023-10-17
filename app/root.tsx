import type { LinksFunction } from '@remix-run/cloudflare'
import { Links, LiveReload, Meta, Outlet, Scripts, ScrollRestoration } from '@remix-run/react'

import stylesheet from '~/tailwind.css'

import Footer from './components/footer'
import Header from './components/header'

export const links: LinksFunction = () => [{ rel: 'stylesheet', href: stylesheet }]

export default function App() {
  return (
    <html lang="en">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body className="m-auto max-w-xl pl-3 pr-3 text-justify">
        <Header />
        <Outlet /> {/* THIS IS WHERE ROUTE PAGES ARE INSERTED */}
        <Footer />
        <ScrollRestoration />
        <Scripts />
        <LiveReload />
      </body>
    </html>
  )
}

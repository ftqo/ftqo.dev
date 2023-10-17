import { Link } from '@remix-run/react'

export default function Header() {
  return (
    <div className="navbar mb-4 flex items-center justify-between">
      <Link to="/" className="text-3xl font-bold">
        ftqo
      </Link>
    </div>
  )
}

import type { MetaFunction } from '@remix-run/cloudflare'

export const meta: MetaFunction = () => {
  return [{ title: 'ftqo site' }, { name: 'description', content: 'WIP' }]
}

export default function Page() {
  return (
    <div className="[&>p]:mb-4">
      <h1 className="mb-4 text-2xl font-bold">Work In Progress</h1>
    </div>
  )
}

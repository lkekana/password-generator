import { createFileRoute, redirect } from '@tanstack/react-router'

export const Route = createFileRoute('/$')({
  beforeLoad: () => {
    throw redirect({
      to: '/',
      replace: true, // Replaces the 404 URL in the history stack
    })
  },
})
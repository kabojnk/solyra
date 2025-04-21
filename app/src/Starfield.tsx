import React, { useEffect, useRef } from 'react'

const STAR_COUNT = 390
const STAR_MIN_SIZE = .3
const STAR_MAX_SIZE = 1
const STARFIELD_OPACITY = 0.8

function randomBetween(min: number, max: number) {
  return Math.random() * (max - min) + min
}

function createStars(width: number, height: number) {
  return Array.from({ length: STAR_COUNT }, () => ({
    x: randomBetween(0, width),
    y: randomBetween(0, height),
    r: randomBetween(STAR_MIN_SIZE, STAR_MAX_SIZE),
    twinkle: randomBetween(0, 1),
  }))
}

export default function Starfield() {
  const canvasRef = useRef<HTMLCanvasElement>(null)

  useEffect(() => {
    const canvas = canvasRef.current
    if (!canvas) return
    const ctx = canvas.getContext('2d')
    if (!ctx) return

    let animationId: number
    let width = window.innerWidth
    let height = window.innerHeight
    canvas.width = width
    canvas.height = height

    let stars = createStars(width, height)

    function draw() {
      
      if (!ctx) return
      
      ctx.clearRect(0, 0, width, height)
      for (const star of stars) {
        const twinkle = 0.7 + 0.3 * Math.sin(Date.now() / 700 + star.twinkle * Math.PI * 2)
        ctx.globalAlpha = twinkle * STARFIELD_OPACITY
        ctx.beginPath()
        ctx.arc(star.x, star.y, star.r, 0, 2 * Math.PI)
        ctx.fillStyle = '#fff'
        ctx.shadowColor = '#fff'
        ctx.shadowBlur = randomBetween(1, 3)
        ctx.fill()
      }
      ctx.globalAlpha = 1
      animationId = requestAnimationFrame(draw)
    }

    draw()

    function handleResize() {
      if (!canvas) return
      
      width = window.innerWidth
      height = window.innerHeight
      canvas.width = width
      canvas.height = height
      stars = createStars(width, height)
    }
    window.addEventListener('resize', handleResize)
    return () => {
      cancelAnimationFrame(animationId)
      window.removeEventListener('resize', handleResize)
    }
  }, [])

  return (
    <canvas
      ref={canvasRef}
      style={{
        position: 'fixed',
        top: 0,
        left: 0,
        width: '100vw',
        height: '100vh',
        zIndex: 0,
        pointerEvents: 'none',
      }}
      aria-hidden="true"
    />
  )
}

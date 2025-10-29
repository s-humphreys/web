import React from 'react'
import { FaLinkedin, FaGithub } from 'react-icons/fa'
import { MdEmail } from 'react-icons/md'
import DividerDots from './DividerDots'

type LinksProps = {
  children?: React.ReactNode // extra <li> items
}

export default function Links({ children }: LinksProps) {
  return (
    <>
      <DividerDots />
      <nav aria-label="Primary">
        <ul className="url-links">
          {children}
          <li>
            <a href="mailto:sam@shumphreys.com">
              <MdEmail className="icon" aria-hidden="true" />
              Email
            </a>
          </li>
          <li>
            <a href="https://www.linkedin.com/in/samgeoffreyhumphreys" target="_blank" rel="noopener noreferrer">
              <FaLinkedin className="icon" aria-hidden="true" />
              LinkedIn
            </a>
          </li>
          <li>
            <a href="https://github.com/s-humphreys" target="_blank" rel="noopener noreferrer">
              <FaGithub className="icon" aria-hidden="true" />
              GitHub
            </a>
          </li>
        </ul>
      </nav>
    </>
  )
}

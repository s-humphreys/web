import React from 'react'
import { Link } from 'react-router-dom'
import { FiFileText } from 'react-icons/fi'
import Links from '../components/Links'

export default function Home() {
  return (
    <main className="container">
      <header>
        <h1>Sam Humphreys</h1>
        <hr className="center-diamond" />
        <h2>
          Senior Platform Engineer •{' '}
          <a href="https://capitalontap.com/" target="_blank" rel="noopener noreferrer">
            Capital on Tap - London, UK
          </a>
        </h2>
      </header>

      <section style={{ marginTop: 8 }}>
        <p>
          Experienced cloud, platform and Kubernetes engineer.
        </p>
        <p>
          Currently focusing on building secure and highly available platforms;
          creating golden paths for engineering success; and implementing
          site reliability engineering principles.
        </p>
        <p>
          Coding mostly in Go, running on GCP/Azure and Kubernetes, leveraging GitOps
          and open source technologies.
        </p>
      </section>

      <Links>
        <li>
          <Link to="/experience">
            <FiFileText className="icon" aria-hidden="true" />
            Experience
          </Link>
        </li>
      </Links>

      <div className="content-push" />
    </main>
  )
}

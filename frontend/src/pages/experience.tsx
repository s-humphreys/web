import React from 'react'
import { Link } from 'react-router-dom'
import { FaHome } from 'react-icons/fa'
import Links from '../components/Links'

type Item = {
  company: string
  url?: string
  role: string
  period: string
  skills: string[]
}

const items: Item[] = [
  {
    company: 'Capital on Tap - London, UK',
    url: 'https://capitalontap.com/',
    role: 'Platform Engineer',
    period: 'July 2024 - Current',
    skills: [
      'Go',
      'Powershell',
      'Kubernetes',
      'Istio',
      'Flux',
      'Terraform',
      'Prometheus',
      'Grafana',
      'DataDog',
      'Azure',
      'ADO',
    ],
  },
  {
    company: 'ESG Book - London, UK',
    url: 'https://www.esgbook.com/',
    role: 'Senior DevOps Engineer',
    period: 'August 2021 - July 2024',
    skills: [
      'Go',
      'Python',
      'Kubernetes',
      'Istio',
      'Vault',
      'Argo',
      'Terraform',
      'Terragrunt',
      'Prometheus',
      'Grafana',
      'GCP',
      'GitHub',
    ],
  },
  {
    company: 'GroupM Data & Technology - London, UK',
    url: 'https://www.groupm.com',
    role: 'Site Reliability Engineer',
    period: 'September 2020 - August 2021',
    skills: [
      'Python',
      'Kubernetes',
      'Istio',
      'Terraform',
      'Cloud Monitoring',
      'GCP',
      'GitLab',
    ],
  },
  {
    company: 'Essence - London, UK',
    url: 'https://www.essencemediacom.com',
    role: 'Site Reliability Engineer',
    period: 'December 2019 - September 2020',
    skills: [
      'Python',
      'Kubernetes',
      'Terraform',
      'Cloud Monitoring',
      'GCP',
      'GitLab',
    ],
  },
  {
    company: 'Dataffirm - London, UK',
    role: 'Software / Infrastructure Engineer',
    period: 'September 2017 - November 2019',
    skills: [
      'Python',
      'Kubernetes',
      'Terraform',
      'Airflow',
      'Cloud Monitoring',
      'GCP',
    ],
  },
]

export default function Experience() {
  return (
    <main className="container">
      <h1>
        <Link className="return-arrow" to="/">&lt;</Link>
        Experience
      </h1>
      <hr className="center-diamond" />

      {items.map((item, i) => (
        <React.Fragment key={item.company + item.period}>
          <div className="work-item">
            <h2>
              {item.url ? (
                <a href={item.url} target="_blank" rel="noopener noreferrer">{item.company}</a>
              ) : (
                item.company
              )}
            </h2>
            <p>{item.role}</p>
            <p className="work-period">{item.period}</p>
            <ul className="skills">
              {item.skills.map(s => <li key={s}>{s}</li>)}
            </ul>
          </div>
          {i < items.length - 1 && <hr className="hr-default" />}
        </React.Fragment>
      ))}

      <Links>
        <li>
          <Link to="/">
            <FaHome className="icon" aria-hidden="true" />
            Home
          </Link>
        </li>
      </Links>

      <div className="content-push" />
    </main>
  )
}

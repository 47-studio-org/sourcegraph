import React, { useEffect } from 'react'

import classNames from 'classnames'

import { Maybe } from '@sourcegraph/search'
import { AuthenticatedUser } from '@sourcegraph/shared/src/auth'
import { ThemeProps } from '@sourcegraph/shared/src/theme'
import { ButtonLink, Container, Link, PageHeader } from '@sourcegraph/wildcard'

import { refreshAuthenticatedUser } from '../../../auth'
import { FeatureFlagProps } from '../../../featureFlags/featureFlags'
import { eventLogger } from '../../../tracking/eventLogger'

import styles from './organizationList.module.scss'
export interface OrganizationsListProps extends ThemeProps, FeatureFlagProps {
    authenticatedUser: Pick<
        AuthenticatedUser,
        'username' | 'avatarURL' | 'settingsURL' | 'organizations' | 'siteAdmin' | 'session' | 'displayName'
    >
}

interface IOrgItem {
    id: string
    name: string
    displayName: Maybe<string>
    url: string
    settingsURL: Maybe<string>
}

interface OrgItemProps {
    org: IOrgItem
}

const OrgItem: React.FunctionComponent<OrgItemProps> = ({ org }) => (
    <li data-test-username={org.id}>
        <div className="d-flex align-items-center justify-content-between">
            <div className={classNames('d-flex align-items-center justify-content-start flex-1', styles.orgDetails)}>
                <div className={styles.avatarContainer}>
                    <div className={styles.avatar}>
                        <span>{(org.displayName || org.name).slice(0, 2).toUpperCase()}</span>
                    </div>
                </div>
                <div className="d-flex flex-column">
                    <Link to={org.url} className={styles.orgLink}>
                        {org.displayName || org.name}
                    </Link>
                    {org.displayName && (
                        <span className={classNames('text-muted', styles.displayName)}>{org.name}</span>
                    )}
                </div>
            </div>

            <div className={styles.userRole}>
                <span className="text-muted">Admin</span>
            </div>
            <div>
                <ButtonLink className={styles.orgSettings} variant="secondary" to={org.settingsURL as string} size="sm">
                    Settings
                </ButtonLink>
            </div>
        </div>
    </li>
)

const refreshOrganizationList = (): void => {
    refreshAuthenticatedUser()
        .toPromise()
        .then(() => {
            eventLogger.logViewEvent('OrganizationsList')
        })
        .catch(() => eventLogger.logViewEvent('ErrorOrgListLoading'))
}

export const OrganizationsListPage: React.FunctionComponent<OrganizationsListProps> = ({ authenticatedUser }) => {
    useEffect(() => {
        refreshOrganizationList()
    }, [])

    const orgs = authenticatedUser.organizations.nodes
    const hasOrgs = orgs.length > 0

    return (
        <div className="org-list-page">
            <div className="d-flex flex-0 justify-content-end align-items-center mb-3 flex-wrap">
                <PageHeader path={[{ text: 'Organizations' }]} headingElement="h2" className="flex-1" />

                <ButtonLink variant="secondary" to="/organizations/joinopenbeta" size="sm">
                    Create organization
                </ButtonLink>
            </div>
            {hasOrgs && (
                <Container className={classNames('mb-4', styles.organisationsList)}>
                    <ul>
                        {orgs.map(org => (
                            <OrgItem org={org} key={org.id} />
                        ))}
                    </ul>
                </Container>
            )}
            {!hasOrgs && (
                <Container>
                    <div className="d-flex flex-0 flex-column justify-content-center align-items-center">
                        <h3>Start searching with your team on Sourcegraph</h3>
                        <div>Product copy here that needs to be written still, this is a placeholder.</div>
                        <ButtonLink variant="primary" to="/organizations/joinopenbeta" size="sm" className="mt-3">
                            Create your organization
                        </ButtonLink>
                    </div>
                </Container>
            )}
        </div>
    )
}

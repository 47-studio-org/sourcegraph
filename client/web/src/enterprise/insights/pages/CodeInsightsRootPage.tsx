import React, { useEffect } from 'react'

import PlusIcon from 'mdi-react/PlusIcon'
import { matchPath, useHistory } from 'react-router'
import { useLocation } from 'react-router-dom'

import { useTemporarySetting } from '@sourcegraph/shared/src/settings/temporary/useTemporarySetting'
import { TelemetryProps } from '@sourcegraph/shared/src/telemetry/telemetryService'
import { lazyComponent } from '@sourcegraph/shared/src/util/lazyComponent'
import { Button, Link, PageHeader, Tabs, TabList, Tab, Icon } from '@sourcegraph/wildcard'

import { CodeInsightsIcon } from '../../../insights/Icons'
import { CodeInsightsPage } from '../components/code-insights-page/CodeInsightsPage'
import { ALL_INSIGHTS_DASHBOARD_ID } from '../core/types/dashboard/virtual-dashboard'

import { DashboardsContentPage } from './dashboards/dashboard-page/DashboardsContentPage'

const LazyCodeInsightsGettingStartedPage = lazyComponent(
    () => import('./landing/getting-started/CodeInsightsGettingStartedPage'),
    'CodeInsightsGettingStartedPage'
)

export enum CodeInsightsRootPageURLPaths {
    CodeInsights = '/dashboards/:dashboardId?',
    GettingStarted = '/about',
}

export enum CodeInsightsRootPageTab {
    CodeInsights,
    GettingStarted,
}

function useQuery(): URLSearchParams {
    const { search } = useLocation()

    return React.useMemo(() => new URLSearchParams(search), [search])
}

interface CodeInsightsRootPageProps extends TelemetryProps {
    activeView: CodeInsightsRootPageTab
}

export const CodeInsightsRootPage: React.FunctionComponent<CodeInsightsRootPageProps> = props => {
    const { telemetryService, activeView } = props
    const location = useLocation()
    const query = useQuery()
    const history = useHistory()

    const { params } =
        matchPath<{ dashboardId?: string }>(location.pathname, {
            path: `/insights${CodeInsightsRootPageURLPaths.CodeInsights}`,
        }) ?? {}

    const [hasInsightPageBeenViewed, markMainPageAsViewed] = useTemporarySetting('insights.wasMainPageOpen', false)

    const dashboardId = params?.dashboardId ?? ALL_INSIGHTS_DASHBOARD_ID
    const queryParameterDashboardId = query.get('dashboardId') ?? ALL_INSIGHTS_DASHBOARD_ID

    const handleTabNavigationChange = (selectedTab: CodeInsightsRootPageTab): void => {
        switch (selectedTab) {
            case CodeInsightsRootPageTab.CodeInsights:
                return history.push(`/insights/dashboards/${queryParameterDashboardId}`)
            case CodeInsightsRootPageTab.GettingStarted:
                return history.push(`/insights/about?dashboardId=${dashboardId}`)
        }
    }

    useEffect(() => {
        if (hasInsightPageBeenViewed === false) {
            markMainPageAsViewed(true)
        }
    }, [hasInsightPageBeenViewed, markMainPageAsViewed])

    return (
        <CodeInsightsPage>
            <PageHeader
                path={[{ icon: CodeInsightsIcon }, { text: 'Insights' }]}
                actions={
                    <>
                        <Button as={Link} to="/insights/add-dashboard" variant="secondary" className="mr-2">
                            <Icon as={PlusIcon} /> Add dashboard
                        </Button>
                        <Button
                            as={Link}
                            to={`/insights/create?dashboardId=${dashboardId}`}
                            variant="primary"
                            onClick={() => telemetryService.log('InsightAddMoreClick')}
                        >
                            <Icon as={PlusIcon} /> Create insight
                        </Button>
                    </>
                }
                className="align-items-start mb-3"
            />

            <Tabs index={activeView} size="medium" className="mb-3" onChange={handleTabNavigationChange}>
                <TabList>
                    <Tab index={CodeInsightsRootPageTab.CodeInsights}>Code Insights</Tab>

                    <Tab index={CodeInsightsRootPageTab.GettingStarted}>Getting started</Tab>
                </TabList>
            </Tabs>

            {activeView === CodeInsightsRootPageTab.CodeInsights && (
                <DashboardsContentPage telemetryService={telemetryService} dashboardID={params?.dashboardId} />
            )}

            {activeView === CodeInsightsRootPageTab.GettingStarted && (
                <LazyCodeInsightsGettingStartedPage telemetryService={telemetryService} />
            )}
        </CodeInsightsPage>
    )
}

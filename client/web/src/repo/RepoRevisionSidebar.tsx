import React, { useCallback, useState } from 'react'

import classNames from 'classnames'
import * as H from 'history'
import ChevronDoubleLeftIcon from 'mdi-react/ChevronDoubleLeftIcon'
import ChevronDoubleRightIcon from 'mdi-react/ChevronDoubleRightIcon'

import { Resizable } from '@sourcegraph/shared/src/components/Resizable'
import { ExtensionsControllerProps } from '@sourcegraph/shared/src/extensions/controller'
import { Scalars } from '@sourcegraph/shared/src/graphql-operations'
import { TelemetryProps } from '@sourcegraph/shared/src/telemetry/telemetryService'
import { ThemeProps } from '@sourcegraph/shared/src/theme'
import { AbsoluteRepoFile } from '@sourcegraph/shared/src/util/url'
import { Button, useLocalStorage, useMatchMedia, Tab, TabList, TabPanel, TabPanels, Tabs } from '@sourcegraph/wildcard'

import settingsSchemaJSON from '../../../../schema/settings.schema.json'
import { GettingStartedTour } from '../gettingStartedTour/GettingStartedTour'
import { Tree } from '../tree/Tree'

import { RepoRevisionSidebarSymbols } from './RepoRevisionSidebarSymbols'

import styles from './RepoRevisionSidebar.module.scss'

interface Props extends AbsoluteRepoFile, ExtensionsControllerProps, ThemeProps, TelemetryProps {
    repoID: Scalars['ID']
    isDir: boolean
    defaultBranch: string
    className: string
    history: H.History
    location: H.Location
    showGettingStartedTour?: boolean
}

const SIZE_STORAGE_KEY = 'repo-revision-sidebar'
const TABS_KEY = 'repo-revision-sidebar-last-tab'
const SIDEBAR_KEY = 'repo-revision-sidebar-toggle'
/**
 * The sidebar for a specific repo revision that shows the list of files and directories.
 */
export const RepoRevisionSidebar: React.FunctionComponent<Props> = props => {
    const [persistedTabIndex, setPersistedTabIndex] = useLocalStorage(TABS_KEY, 0)
    const [persistedIsVisible, setPersistedIsVisible] = useLocalStorage(
        SIDEBAR_KEY,
        settingsSchemaJSON.properties.fileSidebarVisibleByDefault.default
    )

    const isWideScreen = useMatchMedia('(min-width: 768px)', false)
    const [isVisible, setIsVisible] = useState(persistedIsVisible && isWideScreen)

    const handleSidebarToggle = useCallback(
        (value: boolean) => {
            props.telemetryService.log('FileTreeViewClicked', {
                action: 'click',
                label: 'expand / collapse file tree view',
            })
            setPersistedIsVisible(value)
            setIsVisible(value)
        },
        [setPersistedIsVisible, props.telemetryService]
    )
    const handleSymbolClick = useCallback(() => props.telemetryService.log('SymbolTreeViewClicked'), [
        props.telemetryService,
    ])

    if (!isVisible) {
        return (
            <Button
                variant="icon"
                className={classNames('position-absolute border-top border-bottom border-right mt-4', styles.toggle)}
                onClick={() => handleSidebarToggle(true)}
                data-tooltip="Show sidebar"
            >
                <ChevronDoubleRightIcon className="icon-inline" />
            </Button>
        )
    }

    return (
        <Resizable
            defaultSize={256}
            handlePosition="right"
            storageKey={SIZE_STORAGE_KEY}
            element={
                <div className="d-flex flex-column w-100">
                    {props.showGettingStartedTour && (
                        <GettingStartedTour className="mb-1 mr-3" telemetryService={props.telemetryService} />
                    )}
                    <Tabs
                        className="w-100 h-100 test-repo-revision-sidebar pr-3"
                        defaultIndex={persistedTabIndex}
                        onChange={setPersistedTabIndex}
                        lazy={true}
                    >
                        <TabList
                            actions={
                                <Button
                                    onClick={() => handleSidebarToggle(false)}
                                    className="bg-transparent border-0 ml-auto p-1 position-relative focus-behaviour"
                                    title="Hide sidebar"
                                    data-tooltip="Hide sidebar"
                                    data-placement="right"
                                >
                                    <ChevronDoubleLeftIcon className={classNames('icon-inline', styles.closeIcon)} />
                                </Button>
                            }
                        >
                            <Tab data-tab-content="files">
                                <span className="tablist-wrapper--tab-label">Files</span>
                            </Tab>
                            <Tab data-tab-content="symbols">
                                <span className="tablist-wrapper--tab-label">Symbols</span>
                            </Tab>
                        </TabList>
                        <div
                            aria-hidden={true}
                            className={classNames('flex w-100 overflow-auto explorer', styles.tabpanels)}
                            tabIndex={-1}
                        >
                            <TabPanels>
                                <TabPanel>
                                    <Tree
                                        key="files"
                                        repoName={props.repoName}
                                        revision={props.revision}
                                        commitID={props.commitID}
                                        history={props.history}
                                        location={props.location}
                                        scrollRootSelector=".explorer"
                                        activePath={props.filePath}
                                        activePathIsDir={props.isDir}
                                        sizeKey={`Resizable:${SIZE_STORAGE_KEY}`}
                                        extensionsController={props.extensionsController}
                                        isLightTheme={props.isLightTheme}
                                        telemetryService={props.telemetryService}
                                    />
                                </TabPanel>
                                <TabPanel>
                                    <RepoRevisionSidebarSymbols
                                        key="symbols"
                                        repoID={props.repoID}
                                        revision={props.revision}
                                        activePath={props.filePath}
                                        onHandleSymbolClick={handleSymbolClick}
                                    />
                                </TabPanel>
                            </TabPanels>
                        </div>
                    </Tabs>
                </div>
            }
        />
    )
}

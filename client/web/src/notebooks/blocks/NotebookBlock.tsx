import React, { useEffect, useMemo } from 'react'

import classNames from 'classnames'

import { isMacPlatform as isMacPlatformFn } from '@sourcegraph/common'

import { BlockProps } from '..'
import { isModifierKeyPressed } from '../notebook/useNotebookEventHandlers'

import { NotebookBlockMenu, NotebookBlockMenuProps } from './menu/NotebookBlockMenu'
import { useIsBlockInputFocused } from './useIsBlockInputFocused'

import blockStyles from './NotebookBlock.module.scss'

interface NotebookBlockProps extends Pick<BlockProps, 'isSelected' | 'isOtherBlockSelected'>, NotebookBlockMenuProps {
    className?: string
    'aria-label': string
    onDoubleClick?: () => void
    onEnterBlock: () => void
    onHideInput?: () => void
}

export const NotebookBlock: React.FunctionComponent<NotebookBlockProps> = ({
    children,
    id,
    className,
    isSelected,
    isOtherBlockSelected,
    mainAction,
    actions,
    'aria-label': ariaLabel,
    onDoubleClick,
    onEnterBlock,
    onHideInput,
}) => {
    const isInputFocused = useIsBlockInputFocused(id)
    const isMacPlatform = useMemo(() => isMacPlatformFn(), [])
    useEffect(() => {
        const handleKeyDown = (event: KeyboardEvent): void => {
            if (isSelected && event.key === 'Enter') {
                if (isModifierKeyPressed(event.metaKey, event.ctrlKey, isMacPlatform)) {
                    onHideInput?.()
                } else {
                    onEnterBlock()
                }
            }
        }

        document.addEventListener('keydown', handleKeyDown)
        return () => {
            document.removeEventListener('keydown', handleKeyDown)
        }
    }, [isMacPlatform, isSelected, onEnterBlock, onHideInput])

    return (
        <div className={classNames('block-wrapper', blockStyles.blockWrapper)} data-block-id={id}>
            {/* Notebook blocks are a form of specialized UI for which there are no good accesibility settings (role, aria-*)
            or semantic elements that would accurately describe its functionality. To provide the necessary functionality we have
            to rely on plain div elements and custom click/focus/keyDown handlers. We still preserve the ability to navigate through blocks
            with the keyboard using the up and down arrows, and TAB. */}
            <div
                className={classNames(
                    'block',
                    blockStyles.block,
                    className,
                    isSelected && !isInputFocused && blockStyles.selected,
                    isSelected && isInputFocused && blockStyles.selectedNotFocused
                )}
                onDoubleClick={onDoubleClick}
                // A tabIndex is necessary to make the block focusable.
                // eslint-disable-next-line jsx-a11y/no-noninteractive-tabindex
                tabIndex={0}
                aria-label={ariaLabel}
            >
                {children}
            </div>
            {(isSelected || !isOtherBlockSelected) && (
                <NotebookBlockMenu id={id} mainAction={mainAction} actions={actions} />
            )}
        </div>
    )
}

// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useEffect, useState } from 'react';
import { useDispatch } from 'react-redux';

import { PaginatorBlockPages } from '@components/common/Paginator/PaginatorBlockPages';

import { ArrowNextIcon, ArrowPrevIcon } from '@/app/static/img/UltimateDivisionPaginator/arrows';

import './index.scss';

export const Paginator: React.FC<{
    getCardsOnPage: (selectedPage: number) => void;
    itemsCount: number;
    selectedPage: number;
}> = ({ getCardsOnPage, itemsCount, selectedPage }) => {
    const dispatch = useDispatch();
    const [currentPage, setCurrentPage] = useState<number>(selectedPage);

    /**
     * split the page into 3 blocks that can be needed
     * to separate page numbers
     */
    const [firstBlockPages, setFirstBlockPages] = useState<number[]>([]);
    const [middleBlockPages, setMiddleBlockPages] = useState<number[]>([]);
    const [lastBlockPages, setLastBlockPages] = useState<number[]>([]);

    const CARDS_ON_PAGE: number = 24;
    const MAX_PAGES_PER_BLOCK: number = 5;
    const MAX_PAGES_OFF_BLOCKS: number = 10;
    const FIRST_PAGE_INDEX: number = 0;
    const SECOND_PAGE_INDEX: number = 1;
    const FIRST_PAGE_INDEX_FROM_END: number = -1;
    const NEG_STEP_FROM_CURRENT_PAGE: number = -3;
    const POS_STEP_FROM_CURRENT_PAGE: number = 2;
    const FIRST_PAGE: number = 1;

    /** inplements opening current cards on selected page */
    async function getCards(selectedPage: number) {
        try {
            /** dispatch getCardsOnPage thunk with parameters: page and default limit value */
            await dispatch(
                getCardsOnPage(selectedPage)
            );
        } catch (error: any) {
            /** TODO: it will be reworked with notification system */
        }
    }

    const pages: number[] = [];

    /** Couns pages number. */
    const pagesCount = Math.ceil(itemsCount / CARDS_ON_PAGE);

    for (let i = 1; i <= pagesCount; i++) {
        pages.push(i);
    }

    /**
     * indicates if current page is first page or last page
     */
    const isFirstPageSelected: boolean = currentPage === FIRST_PAGE;
    const isLastPageSelected: boolean = currentPage === pagesCount;

    /** set block pages depends on current page */
    const setBlocksIfCurrentInFirstBlock = () => {
        setFirstBlockPages(pages.slice(FIRST_PAGE_INDEX, MAX_PAGES_PER_BLOCK));
        setMiddleBlockPages([]);
        setLastBlockPages(pages.slice(FIRST_PAGE_INDEX_FROM_END));
    };
    const setBlocksIfCurrentInMiddleBlock = () => {
        setFirstBlockPages(pages.slice(FIRST_PAGE_INDEX, SECOND_PAGE_INDEX));
        setMiddleBlockPages(
            pages.slice(
                currentPage + NEG_STEP_FROM_CURRENT_PAGE,
                currentPage + POS_STEP_FROM_CURRENT_PAGE
            )
        );
        setLastBlockPages(pages.slice(FIRST_PAGE_INDEX_FROM_END));
    };
    const setBlocksIfCurrentInLastBlock = () => {
        setFirstBlockPages(pages.slice(FIRST_PAGE_INDEX, SECOND_PAGE_INDEX));
        setMiddleBlockPages([]);
        setLastBlockPages(pages.slice(-MAX_PAGES_PER_BLOCK));
    };

    /**
     * Indicates visibility of dots after first pages block
     */
    const isFirstDotsShown: boolean =
        middleBlockPages.length <= MAX_PAGES_PER_BLOCK &&
        pages.length > MAX_PAGES_OFF_BLOCKS;
    /*
     * Indicates visibility of dots after middle pages block
     */
    const isSecondDotsShown: boolean = !!middleBlockPages.length;
    /**
     * indicates in which block current page
     */
    const isCurrentInFirstBlock: boolean = currentPage < MAX_PAGES_PER_BLOCK;
    const isCurrentInLastBlock: boolean =
        pages.length - currentPage < MAX_PAGES_PER_BLOCK - SECOND_PAGE_INDEX;
    /**
     * change page blocks reorganization depends
     * on current page
     */
    const isOneBlockRequired: boolean = pages.length <= MAX_PAGES_OFF_BLOCKS;

    const reorganizePagesBlock = () => {
        if (isOneBlockRequired) {
            return;
        }

        if (isCurrentInFirstBlock) {
            setBlocksIfCurrentInFirstBlock();

            return;
        }

        if (!isCurrentInFirstBlock && !isCurrentInLastBlock) {
            setBlocksIfCurrentInMiddleBlock();

            return;
        }

        if (isCurrentInLastBlock) {
            setBlocksIfCurrentInLastBlock();
        }
    };

    /*
     * indicates if dots delimiter is needed
     * to separate page numbers
     */
    const populatePages = () => {
        if (!pages.length) {
            return;
        }
        if (isOneBlockRequired) {
            setFirstBlockPages(pages.slice());
            setMiddleBlockPages([]);
            setLastBlockPages([]);

            return;
        }
        reorganizePagesBlock();
    };

    useEffect(() => {
        getCards(currentPage);
        populatePages();
    }, [currentPage, pagesCount]);
    /**
     * change current page and set pages block
     */
    const onPageChange = (
        type: string,
        pageNumber: number = currentPage
    ): void => {
        const STEP_FROM_CURRENT_PAGE = 1;

        switch (type) {
        case 'next page':
            if (pageNumber < pages.length) {
                setCurrentPage(pageNumber + STEP_FROM_CURRENT_PAGE);
            }
            populatePages();

            return;
        case 'previous page':
            if (pageNumber > SECOND_PAGE_INDEX) {
                setCurrentPage(pageNumber - STEP_FROM_CURRENT_PAGE);
            }
            populatePages();

            return;
        case 'change page':
            setCurrentPage(pageNumber);
            populatePages();

            return;
        default:
            populatePages();
        }
    };

    return (
        <section className="ultimatedivision-paginator">
            <div className="ultimatedivision-paginator__wrapper">
                {isFirstPageSelected ?
                    <a className="ultimatedivision-paginator__previous-not-active">
                        <ArrowPrevIcon/>
                        <p className="ultimatedivision-paginator__previous-not-active__title">
                            Previous page
                        </p>
                    </a>
                    :
                    <a
                        className="ultimatedivision-paginator__previous"
                        onClick={() => onPageChange('previous page')}
                    >
                        <ArrowPrevIcon/>
                        <p className="ultimatedivision-paginator__previous__title">
                            Previous page
                        </p>
                    </a>
                }
                <PaginatorBlockPages
                    blockPages={firstBlockPages}
                    onPageChange={onPageChange}
                    currentPage={currentPage}
                />
                {isFirstDotsShown &&
                    <span className="ultimatedivision-paginator__pages__dots">
                        ...
                    </span>
                }
                <PaginatorBlockPages
                    blockPages={middleBlockPages}
                    onPageChange={onPageChange}
                    currentPage={currentPage}
                />
                {isSecondDotsShown &&
                    <span className="ultimatedivision-paginator__pages__dots">
                        ...
                    </span>
                }
                <PaginatorBlockPages
                    blockPages={lastBlockPages}
                    onPageChange={onPageChange}
                    currentPage={currentPage}
                />
                {isLastPageSelected ?
                    <a
                        className="ultimatedivision-paginator__next-not-active"
                    >
                        <p className="ultimatedivision-paginator__next__title">
                            Next page
                        </p>
                        <ArrowNextIcon/>
                    </a>
                    :
                    <a
                        className="ultimatedivision-paginator__next"
                        onClick={() => onPageChange('next page')}
                    >
                        <p className="ultimatedivision-paginator__next__title">
                            Next page
                        </p>
                        <ArrowNextIcon/>
                    </a>
                }
            </div>
        </section>
    );
};

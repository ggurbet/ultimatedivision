//Copyright (C) 2021 Creditor Corp. Group.
//See LICENSE for copying information.
/* eslint-disable */
import { useEffect, useState } from 'react';

import next from '@static/img/UltimateDivisionPaginator/next.png';
import previous from '@static/img/UltimateDivisionPaginator/previous.png';
import { PaginatorBlockPages } from '@components/Paginator/PaginatorBlockPages';

import './index.scss';

export const Paginator: React.FC<{ itemCount: number }> = ({ itemCount }) => {
    const FIRST_ITEM_PAGINATON = 1;
    const [currentPage, setCurrentPage] = useState<number>(FIRST_ITEM_PAGINATON);
    /**
    * split the page into 3 blocks that can be needed
    * to separate page numbers
     */
    const [firstBlockPages, setFirstBlockPages] = useState<number[]>([]);
    const [middleBlockPages, setMiddleBlockPages] = useState<number[]>([]);
    const [lastBlockPages, setLastBlockPages] = useState<number[]>([]);

    useEffect(() => {
        populatePages();
    }, [currentPage]);

    const CARDS_ON_PAGE: number = 8;
    const MAX_PAGES_PER_BLOCK: number = 5;
    const MAX_PAGES_OFF_BLOCKS: number = 10;
    const FIRST_PAGE_INDEX: number = 0;
    const SECOND_PAGE_INDEX: number = 1;
    const FIRST_PAGE_INDEX_FROM_END: number = -1;
    const NEG_STEP_FROM_CURRENT_PAGE: number = -3;
    const POS_STEP_FROM_CURRENT_PAGE: number = 2;

    const pages: number[] = [];
    for (let i = 1; i <= Math.ceil(itemCount / CARDS_ON_PAGE); i++) {
        pages.push(i);
    }
    /** set block pages depends on current page */
    const setBlocksIfCurrentInFirstBlock = () => {
        setFirstBlockPages(pages.slice(FIRST_PAGE_INDEX, MAX_PAGES_PER_BLOCK));
        setMiddleBlockPages([]);
        setLastBlockPages(pages.slice(FIRST_PAGE_INDEX_FROM_END));
    };
    const setBlocksIfCurrentInMiddleBlock = () => {
        setFirstBlockPages(pages.slice(FIRST_PAGE_INDEX, SECOND_PAGE_INDEX));
        setMiddleBlockPages(pages.slice(currentPage + NEG_STEP_FROM_CURRENT_PAGE, currentPage + POS_STEP_FROM_CURRENT_PAGE));
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
        middleBlockPages.length <= MAX_PAGES_PER_BLOCK
        && pages.length > MAX_PAGES_OFF_BLOCKS;
    /*
    * Indicates visibility of dots after middle pages block
     */
    const isSecondDotsShown: boolean = !!middleBlockPages.length;
    /**
     * indicates in which block current page
     */
    const isCurrentInFirstBlock: boolean = currentPage < MAX_PAGES_PER_BLOCK;
    const isCurrentInLastBlock: boolean = pages.length - currentPage < MAX_PAGES_PER_BLOCK - SECOND_PAGE_INDEX;
    /**
     * change page blocks reorganization depends
     * on current page
     */
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
    const isOneBlockRequired: boolean = pages.length <= MAX_PAGES_OFF_BLOCKS;
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
    /**
     * change current page and set pages block
     */
    const onPageChange = (type: string, pageNumber: number = currentPage): void => {
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
                <a className="ultimatedivision-paginator__previous"
                    onClick={() => onPageChange('previous page')}>
                    <img className="ultimatedivision-paginator__previous__arrow"
                        src={previous}
                        alt="Previous page" />
                    <p className="ultimatedivision-paginator__previous__title">
                        Previous page
                    </p>
                </a>
                <PaginatorBlockPages
                    blockPages={firstBlockPages}
                    onPageChange={onPageChange}
                />
                {isFirstDotsShown
                    && <span className="ultimatedivision-paginator__pages__dots">
                        ...</span>}
                <PaginatorBlockPages
                    blockPages={middleBlockPages}
                    onPageChange={onPageChange}
                />
                {isSecondDotsShown
                    && <span className="ultimatedivision-paginator__pages__dots">
                        ...</span>}
                <PaginatorBlockPages
                    blockPages={lastBlockPages}
                    onPageChange={onPageChange}
                />
                <a className="ultimatedivision-paginator__next"
                    onClick={() => onPageChange('next page')}>
                    <p className="ultimatedivision-paginator__next__title">
                        Next page
                    </p>
                    <img className="ultimatedivision-paginator__next__arrow"
                        src={next}
                        alt="Next page" />
                </a>
            </div>
        </section>
    );
};

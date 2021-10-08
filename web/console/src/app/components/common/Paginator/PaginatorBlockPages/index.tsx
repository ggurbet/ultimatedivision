// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

export const PaginatorBlockPages: React.FC<{
    blockPages: number[];
    onPageChange: (type: string, pageNumber?: number) => void;
    currentPage: number;
}> = ({ blockPages, onPageChange, currentPage }) =>
    <ul className="ultimatedivision-paginator__pages">
        {blockPages.map((page, index) =>
            <li
                className={`ultimatedivision-paginator__pages__item${currentPage === page ? '-active' : ''}`}
                key={index}
                onClick={() => onPageChange('change page', page)}
            >
                {page}
            </li>,
        )}
    </ul>;

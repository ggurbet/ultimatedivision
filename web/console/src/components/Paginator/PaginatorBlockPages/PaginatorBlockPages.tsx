/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
 */

import React from 'react';

export const PaginatorBlockPages: React.FC<{
    blockPages: number[];
    onPageChange: (type: string, pageNumber?: number) => void;
}> = ({ blockPages, onPageChange }) =>
    <ul className="ultimatedivision-paginator__pages">
        {blockPages.map((page, index) =>
            <li
                className="ultimatedivision-paginator__pages__item"
                key={index}
                onClick={() => onPageChange('change page', page)}
            >
                {page}
            </li>
        )}
    </ul>
    ;

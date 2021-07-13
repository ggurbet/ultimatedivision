import React from 'react';

export const PaginatorBlockPages: React.FC<{
    blockPages: number[],
    onPageChange: (type: string, pageNumber?: number) => void
}> = ({ blockPages, onPageChange }) => {
    return (
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
    )
};

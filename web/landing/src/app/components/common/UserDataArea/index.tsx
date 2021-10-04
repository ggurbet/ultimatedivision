// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { SetStateAction } from 'react';

export const UserDataArea: React.FC<{
    value: string,
    placeHolder: string,
    handleChange: any,
    className: string,
    type: string,
    error: SetStateAction<string | null>,
    clearError: any,
}> = ({
    value,
    placeHolder,
    handleChange,
    className,
    type,
    error,
    clearError,
}) => {
    return (
        <div className={`${className}__ wrapper`}>
            <input
                className={error ? `${className}-error` : className}
                value={value}
                placeholder={placeHolder}
                onChange={(e) => {
                    handleChange(e.target.value);
                    clearError(null);
                }}
                type={type}
            />
            {error && <label className={`${className}__error`} htmlFor={value}>
                {error}
            </label>}
        </div>
    );
};

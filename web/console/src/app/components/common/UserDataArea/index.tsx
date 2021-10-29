// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { SetStateAction, useEffect, useState } from 'react';

import { useDebounce } from '@/app/hooks/useDebounce';

export const UserDataArea: React.FC<{
    value: string,
    placeHolder: string,
    onChange: any,
    className: string,
    type: string,
    error: SetStateAction<string | null>,
    clearError: any,
    validate: (value: string) => boolean,
}> = ({
    value,
    placeHolder,
    onChange,
    className,
    type,
    error,
    clearError,
    validate,
}) => {
    const DELAY: number = 500;
    /**
    * The value string from input returned by the useDebounce method after 500 milliseconds.
    */
    const debouncedValue: string = useDebounce(value, DELAY);

    /** inline styles for valid input field */
    const [successLabelClassName, setSuccessLabelClassName] =
        useState<string>('');

    useEffect(() => {
        if (!validate(debouncedValue)) {
            setSuccessLabelClassName('');
        } else {
            setSuccessLabelClassName('-check');
        };
    }, [debouncedValue]);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        onChange(e.target.value);
        clearError(null);
    };

    return (
        <div className={`${className}__ wrapper`}>
            <input
                className={error ? `${className}-error` : `${className}${successLabelClassName}`}
                value={value}
                placeholder={placeHolder}
                onChange={handleChange}
                type={type}
            />
            {error && <label className={`${className}__error`} htmlFor={value}>
                {error}
            </label>}
        </div>
    );
};

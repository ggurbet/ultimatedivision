// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/** Hook for get or set value from/to local storage. */
export const useLocalStorage = () => {
    /* Set value to localStorage */
    const setLocalStorageItem = (item: string, value: boolean) => {
        window.localStorage.setItem(item, JSON.stringify(value));
    };

    /* Get value from localStorage */
    const getLocalStorageItem = (item: string) =>
        window.localStorage.getItem(item);

    return [setLocalStorageItem, getLocalStorageItem];
};

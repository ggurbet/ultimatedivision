// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

const content = document.querySelector('.page');

/** Sets or unsets scrolling page. */
export const setScrollAble = (isSettingScroll: boolean) => {
    isSettingScroll ? content?.classList.remove('scroll-unset') : content?.classList.add('scroll-unset');
};

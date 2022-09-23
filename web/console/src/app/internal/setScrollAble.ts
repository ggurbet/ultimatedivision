// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

/** Sets or unsets scrolling page. */
export const setScrollAble = () => {
    const content = document.querySelector('.page');

    if (content?.classList.contains('scroll-unset')) {
        content?.classList.remove('scroll-unset');
    } else {
        content?.classList.add('scroll-unset');
    }
};

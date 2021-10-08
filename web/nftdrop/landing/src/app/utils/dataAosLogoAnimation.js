// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/*eslint-disable*/
export const dataAosLogoAnimation = (id) => {
    switch (id) {
        case 0:
        case 4:
            return 'fade-right';
        case 3:
        case 7:
            return 'fade-left';
        case 1:
        case 2:
            return 'fade-down';
        case 5:
        case 6:
            return 'fade-up';
    }
};

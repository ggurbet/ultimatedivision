// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

/*eslint-disable*/
export const aosDelayLogoAnimation = (id) => {
    switch (id) {
        case 0:
        case 3:
            return 800;
        case 4:
        case 7:
            return 1000;
        case 1:
        case 2:
            return 400;
        case 5:
        case 6:
            return 600;
    }
};


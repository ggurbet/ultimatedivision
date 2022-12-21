// Copyright (C) 2022 Creditor Corp. Group.
// See LICENSE for copying information.

import { toast } from 'react-toastify';

import toastNoficationsMessages from '@/app/configs/toastNoficationsMessages.json';

/** Code which indicates that requests accounts is already processing. */
const RPC_ERROR_CODE = -32002;
/** Error code indicates that card is already minted. */
const ALREADY_MINTED_ERROR_CODE = -32603;
/** Error code indicates that transaction is denied. */
const DENIED_TRANSACTION_CODE = 4001;
/** Indicates that user has insufficient funds. */
const NOT_ENOUGH_ETH_CODE = 'INSUFFICIENT_FUNDS';
/** Indicates that was choosen not correct network. */
const NETWORK_ERROR = 'NETWORK_ERROR';

/** Indicates that was insufficient balance in account */
const CASPER_INSUFFICIENT_BALANCE_CODE = -32008;

/** Toast notifications themes.
 * Contains colored, light and dark themes.
 */
export enum ToastNotificationsThemes {
    colored = 'colored',
    light = 'light',
    dark = 'dark',
};

/** Toast notifications types.
 * I.e, error, info, success, warning. */
export enum ToastNofiticationsTypes {
    error = 'error',
    info = 'info',
    success = 'success',
    warning = 'warning',
};

/** Defines toast notifications with message, toast type and theme. */
export class ToastNotifications {
    /** Notifies user.
    * As default type uses error type, and default theme is colored. */
    static notify(
        message: string,
        type: ToastNofiticationsTypes = ToastNofiticationsTypes.error,
        theme: ToastNotificationsThemes = ToastNotificationsThemes.colored
    ) {
        toast[type](
            message,
            {
                position: toast.POSITION.TOP_RIGHT,
                theme,
            }
        );
    };

    /** Notifies that something wents wrong agreement. */
    static somethingWentsWrong() {
        this.notify(toastNoficationsMessages.somethingWentsWrong);
    };

    /** Notifies that could not requests artist subscriptions list. */
    static gameCanceled() {
        this.notify(toastNoficationsMessages.gameCanceled);
    };

    /** Notifies that could not requests user profile. */
    static failedToOpenLootbox() {
        this.notify(toastNoficationsMessages.failedToOpenLootbox);
    };

    /** Notifies that could not logout. */
    static failedGettingSeasonStatistics() {
        this.notify(toastNoficationsMessages.failedGettingSeasonStatistics);
    };

    /** Notifies that could not open card. */
    static couldNotCreateClub() {
        this.notify(toastNoficationsMessages.couldNotCreateClub);
    };

    /** Notifies that could not update place bid. */
    static couldNotPlaceBid() {
        this.notify(toastNoficationsMessages.couldNotPlaceBid);
    };

    /** Notifies that could not update profile picture. */
    static couldNotAddCasperWallet() {
        this.notify(toastNoficationsMessages.couldNotAddCasperWallet);
    };

    /** Notifies that could not updated username. */
    static gameFinished() {
        this.notify(toastNoficationsMessages.gameFinished, ToastNofiticationsTypes.info);
    };

    /** Notifies user that email change code sent. */
    static couldNotLogInUserWithCasper() {
        this.notify(
            toastNoficationsMessages.couldNotLogInUserWithCasper,
        );
    };

    /** Notifies that email is changed. */
    static couldNotLogInUserWithMetamask() {
        this.notify(
            toastNoficationsMessages.couldNotLogInUserWithMetamask,
        );
    };

    /** Notifies that email is already in use. */
    static couldNotGetUser() {
        this.notify(toastNoficationsMessages.couldNotGetUser);
    };

    /** Notifies that email is not correct. */
    static couldNotLogInUserWithVelas() {
        this.notify(toastNoficationsMessages.couldNotLogInUserWithVelas);
    };

    /** Notifies to enter email. */
    static registrationFailed() {
        this.notify(toastNoficationsMessages.registrationFailed);
    };

    /** Handles metamask errors and notifies user. */
    static metamaskError(error: any) {
        let errorMessage = '';
        switch (error.code) {
        case RPC_ERROR_CODE:
            errorMessage = toastNoficationsMessages.openMetamaskManually;
            break;
        case DENIED_TRANSACTION_CODE:
            errorMessage = toastNoficationsMessages.transactionDenied;
            break;
        case NOT_ENOUGH_ETH_CODE:
            errorMessage = toastNoficationsMessages.notEnoughBalance;
            break;
        case NETWORK_ERROR:
            errorMessage = toastNoficationsMessages.networkChanged;
            break;
        case ALREADY_MINTED_ERROR_CODE:
            errorMessage = toastNoficationsMessages.cardIsMinted;
            break;
        default:
            errorMessage = toastNoficationsMessages.tryLater;
            break;
        }
        this.notify(errorMessage);
    };

    /** Handles casper errors and notifies user. */
    static casperError(error: any) {
        let errorMessage = '';
        if (error.includes(CASPER_INSUFFICIENT_BALANCE_CODE)) {
            errorMessage = toastNoficationsMessages.insufficientBalance;
        } else {
            errorMessage = toastNoficationsMessages.tryLater;
        }
        this.notify(errorMessage);
    };

    /** Notifies that action was success. */
    static couldNotOpenCard() {
        this.notify(toastNoficationsMessages.couldNotOpenCard);
    };

    /** Notifies that action was success. */
    static success() {
        this.notify(toastNoficationsMessages.success, ToastNofiticationsTypes.success);
    };

    /** Notifies that transaction was successfull. */
    static successfullTransaction() {
        this.notify(toastNoficationsMessages.successfullTransaction, ToastNofiticationsTypes.success);
    };

    /** Notifies that username has been updated. */
    static notFound() {
        this.notify(toastNoficationsMessages.notFound);
    };
};

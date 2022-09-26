// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { toast, ToastOptions } from 'react-toastify';
import { InternalError, TooManyRequestsError } from '@/api';

/** Code which indicates that 'eth_requestAccounts' already processing */
const RPC_ERROR_CODE = -32002;
const ALREADY_MINTED_ERROR_CODE = -32603;
const DENIED_TRANSACTION_CODE = 4001;

const notificationConfig: ToastOptions<{}> = {
    position: toast.POSITION.TOP_RIGHT,
    theme: 'colored',
};

/** Metamask notifications */
export function metamaskNotifications(error: any) {
    switch (error.code) {
    case RPC_ERROR_CODE:
        toast.error('Please open metamask manually!', notificationConfig);
        break;
    case DENIED_TRANSACTION_CODE:
        toast.error('You denied transaction', notificationConfig);
        break;
    default:
        if (error instanceof TooManyRequestsError || error instanceof InternalError) {
            toast.error(error.message, notificationConfig);
        } else if (error.error.code === ALREADY_MINTED_ERROR_CODE) {
            toast.error('Token already minted', notificationConfig);
        }
    }
}

// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ethers } from 'ethers';

import { Service } from '@/ethers/service';
/** Web3 provider */
export const web3Provider = window.ethereum && new ethers.providers.Web3Provider(window.ethereum);

/** Class for creating ethers service */
export class ServicePlugin {
    /** Creates ethers provider instance */
    public static create() {
        return new Service(web3Provider);
    }
}

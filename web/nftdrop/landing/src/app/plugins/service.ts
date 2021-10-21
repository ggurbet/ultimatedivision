// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { ethers } from 'ethers';

import { Service } from '@/app/ethers/service';

export class ServicePlugin {
    public static create() {
        try {
            //@ts-ignore
            const ethereumProvider = new ethers.providers.Web3Provider(window.ethereum);

            return new Service(ethereumProvider);
        } catch (e) {
            return new Service(null);
        }

    }
}

// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

import { HttpClient } from '@/private/http/client';

/**
 * BadRequestError is a custom error type for performing bad request.
 */
export class BadRequestError extends Error {
    public constructor(message = 'User with this email already exists') {
        super(message);
    }
};

/**
 * InternalError is a custom error type for internal server error.
 */
export class InternalError extends Error {
    public constructor(message = 'Something is wrong. Please, try later') {
        super(message);
    }
};

/**
 * APIClient is base client that holds http client and error handler.
 */
export class APIClient {
    protected readonly http: HttpClient = new HttpClient();
    protected async handleError(response: Response): Promise<void> {
        switch (response.status) {
            case 400: throw new BadRequestError();
            case 500: throw new InternalError();
            default:
                throw new InternalError();
        }
    }
};

// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { useState, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useHistory } from 'react-router-dom';
import { toast } from 'react-toastify';

import { AutoCloseTimer } from './AutoCloseTimer';
import { Timer } from './Timer';

import { QueueClient } from '@/api/queue';
import { RouteConfig } from '@/app/routes';
import { RootState } from '@/app/store';
import { getMatchScore } from '@/app/store/actions/mathes';
import { startSearchingMatch } from '@/app/store/actions/clubs';

import './index.scss';

const MatchFinder: React.FC = () => {
    const { squad } = useSelector(
        (state: RootState) => state.clubsReducer.activeClub
    );
    const { isSearchingMatch } = useSelector(
        (state: RootState) => state.clubsReducer
    );

    /** Indicates that user have rejected game. */
    const [isRejectedUser, setIsRejectedUser] = useState<boolean>(false);

    const [queueClient, setQueueClient] = useState<QueueClient | null>(null);

    const dispatch = useDispatch();
    const history = useHistory();

    /** Indicates if match is found. */
    const [isMatchFound, setIsMatchFound] = useState<boolean>(false);

    /** Delay is time delay for redirect user to match page. */
    const DELAY: number = 2000;

    /** DELAY_AFTER_REJECT is time delay in milliseconds for searching match after reject. */
    const DELAY_AFTER_REJECT: number = 500;

    /** Variables describes first and second teams indexes for eventAction response. */
    const FIRST_TEAM_INDEX: number = 0;
    const SECOND_TEAM_INDEX: number = 1;

    /** Variable describes that webscoket connection responsed with error. */
    const ERROR_MESSAGE: string = 'could not write to websocket';
    /** Variable describes that user still searching game. */
    const STILL_SEARCHING_MESSAGE: string = 'you are still in search!';
    /** Variable describes that was send wrong action from user. */
    const WRONG_ACTION_MESSAGE: string = 'wrong action';
    /** Variable describes that user added to gueue. */
    const YOU_ADDED_MESSAGE: string = 'you added!';
    /** Variable describes that it needs confirm game from user. */
    const YOU_CONFIRM_PLAY_MESSAGE: string = 'you confirm play?';
    /** Variable describes that user have leaved from searching game. */
    const YOU_LEAVED_MESSAGE: string = 'you leaved!';

    /** Exposes confirm match logic. */
    const confirmMatch = () => {
        queueClient && queueClient.sendAction('confirm', squad.id);
    };

    /** Canceles confirmation game. */
    const cancelConfirmationGame = () => {
        queueClient && queueClient.sendAction('reject', squad.id);
        setIsRejectedUser(true);
    };

    // TODO: rework after ./queue/chore.go solution.
    /** Starts searching match after rejected by user. */
    const startSearchAfterReject = () => {
        const newQueueClient = new QueueClient();
        newQueueClient.startSearch('startSearch', squad.id);

        setQueueClient(newQueueClient);

        toast.error('Your game was canceled. You are still in search.', {
            position: toast.POSITION.TOP_RIGHT,
            theme: 'colored',
        });
    };

    /** Canceles searching game and closes MatchFinder component. */
    const canselSearchingGame = () => {
        // TODO: rework after ./queue/chore.go solution
        queueClient && queueClient.ws.close();

        const newQueueClient = new QueueClient();
        newQueueClient.finishSearch('finishSearch', squad.id);

        setQueueClient(newQueueClient);
        dispatch(startSearchingMatch(false));
    };

    /** Exposes start searching match logic. */
    const startSearchMatch = () => {
        // TODO: rework after ./queue/chore.go solution.
        const newQueueClient = new QueueClient();
        newQueueClient.startSearch('startSearch', squad.id);
        setQueueClient(newQueueClient);
    };

    useEffect(() => {
        isSearchingMatch && startSearchMatch();
    }, [isSearchingMatch]);

    /** Processes queue client event messages. */
    if (queueClient) {
        queueClient.ws.onmessage = ({ data }: MessageEvent) => {
            const messageEvent = JSON.parse(data);

            switch (messageEvent.message) {
            case ERROR_MESSAGE:
                toast.error('error message', {
                    position: toast.POSITION.TOP_RIGHT,
                    theme: 'colored',
                });

                return;
            case STILL_SEARCHING_MESSAGE:
                /** TODO: will be deleted after ./queue/chore.go reworks. */
                queueClient && queueClient.ws.close();
                setIsMatchFound(false);

                if (isRejectedUser) {
                    setTimeout(() => {
                        startSearchAfterReject();
                    }, DELAY_AFTER_REJECT);

                    setIsRejectedUser(false);

                    return;
                };

                startSearchAfterReject();

                return;
            case WRONG_ACTION_MESSAGE:
                toast.error('Something wrong, please, try later.', {
                    position: toast.POSITION.TOP_RIGHT,
                    theme: 'colored',
                });

                return;
            case YOU_ADDED_MESSAGE:
                setIsMatchFound(false);

                return;
            case YOU_CONFIRM_PLAY_MESSAGE:
                setIsMatchFound(true);

                return;
            case YOU_LEAVED_MESSAGE:
                dispatch(startSearchingMatch(false));

                return;
            default:
                const teams = messageEvent.message.matchResults;

                const transaction = messageEvent.message.transaction;
                toast.success(
                    'Successfully! You will be redirected to match page',
                    {
                        position: toast.POSITION.TOP_RIGHT,
                    }
                );

                dispatch(getMatchScore({ teams, transaction }));
                dispatch(startSearchingMatch(false));

                /** implements redirect to match page after DELAY time.  */
                setTimeout(() => {
                    history.push(RouteConfig.Match.path);
                }, DELAY);
            }
        };
    }

    if (queueClient) {
        queueClient.ws.onerror = (event: Event) => {
            toast.error('Something wrong, please, try later.', {
                position: toast.POSITION.TOP_RIGHT,
                theme: 'colored',
            });
        };
    }

    return isSearchingMatch && <section className={isMatchFound ? 'match-finder__wrapper' : ''}>
        <div className="match-finder">
            <h1 className="match-finder__title">
                {isMatchFound ? 'YOUR MATCH WAS FOUND' : 'LOOKING FOR A MATCH'}
            </h1>
            {isMatchFound ? <AutoCloseTimer /> : <Timer />}
            <div className="match-finder__form">
                {isMatchFound ? <input
                    className="match-finder__form__accept"
                    value="Accept"
                    type="button"
                    onClick={confirmMatch}
                /> : <input
                    className="match-finder__form__accept-not-allowed"
                    value="Accept"
                    type="button"
                />}
                {isMatchFound ? <input
                    className="match-finder__form__cancel"
                    value="Cancel"
                    type="button"
                    onClick={cancelConfirmationGame}
                /> : <input
                    className="match-finder__form__cancel"
                    value="Cancel"
                    type="button"
                    onClick={canselSearchingGame}
                />}
            </div>
        </div>
    </section>;
};

export default MatchFinder;

// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import { filteredLots } from '@/app/store/actions/cards';

import { useMarketplace } from '@/app/hooks/marketplace';

import { MarketPlaceCardsGroup } from '@components/MarketPlace/MarketPlaceCardsGroup';
import { FilterField } from '@components/common/FilterField';
import { Paginator } from '@components/common/Paginator';

import './index.scss';

const MarketPlace: React.FC = () => {
    /** TODO: decide use custom hook or directly dispatch thunk into useEffect*/
    const lots = useMarketplace();

    return (
        <section className="marketplace">
            <FilterField
                title="MARKETPLACE"
                thunk={filteredLots}
            />
            <MarketPlaceCardsGroup
                lots={lots}
            />
            <Paginator
                itemCount={lots.length} />
        </section>
    );
};

export default MarketPlace;

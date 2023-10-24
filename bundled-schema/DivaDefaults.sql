BEGIN;

-- Ripped prizes
INSERT INTO public.diva_prizes
(type, points_req, item_type, item_id, quantity, gr, repeatable)
VALUES
    ('personal', 1, 7, 13021, 1, false, false),
    ('personal', 1, 7, 13021, 1, true, false),
    ('personal', 200, 7, 1472, 5, false, false),
    ('personal', 200, 7, 7976, 5, true, false),
    ('personal', 400, 26, 0, 500, false, false),
    ('personal', 400, 26, 0, 500, true, false),
    ('personal', 600, 7, 1472, 5, false, false),
    ('personal', 600, 7, 7976, 5, true, false),
    ('personal', 800, 26, 0, 1000, false, false),
    ('personal', 800, 26, 0, 1000, true, false),
    ('personal', 1000, 26, 0, 1200, false, false),
    ('personal', 1000, 26, 0, 1200, true, false),
    ('personal', 1200, 26, 0, 1500, false, false),
    ('personal', 1200, 26, 0, 1500, true, false),
    ('personal', 1400, 26, 0, 2300, false, false),
    ('personal', 1400, 26, 0, 2300, true, false),
    ('personal', 1600, 26, 0, 2500, false, false),
    ('personal', 1600, 26, 0, 2500, true, false),
    ('personal', 1800, 26, 0, 3000, false, false),
    ('personal', 1800, 26, 0, 3000, true, false),
    ('personal', 2000, 7, 9722, 1, false, false),
    ('personal', 2000, 7, 9723, 1, false, false),
    ('personal', 2000, 7, 9724, 1, false, false),
    ('personal', 2000, 26, 0, 3300, false, false),
    ('personal', 2000, 7, 9722, 1, true, false),
    ('personal', 2000, 7, 9723, 1, true, false),
    ('personal', 2000, 7, 9724, 1, true, false),
    ('personal', 2000, 26, 0, 3300, true, false),
    ('personal', 3000, 7, 1472, 5, false, false),
    ('personal', 3000, 7, 7976, 5, true, false),
    ('personal', 4000, 26, 0, 3500, false, false),
    ('personal', 4000, 26, 0, 3500, true, false),
    ('personal', 5000, 7, 1472, 5, false, false),
    ('personal', 5000, 7, 7976, 5, true, false),
    ('personal', 6000, 7, 9725, 1, false, false),
    ('personal', 6000, 7, 9726, 1, false, false),
    ('personal', 6000, 7, 9727, 1, false, false),
    ('personal', 6000, 7, 9725, 1, true, false),
    ('personal', 6000, 7, 9726, 1, true, false),
    ('personal', 6000, 7, 9727, 1, true, false),
    ('personal', 7000, 26, 0, 3700, false, false),
    ('personal', 7000, 26, 0, 3700, true, false),
    ('personal', 8000, 7, 10192, 5, false, false),
    ('personal', 8000, 7, 10192, 5, true, false),
    ('personal', 9000, 26, 0, 4000, false, false),
    ('personal', 9000, 26, 0, 4000, true, false),
    ('personal', 10000, 7, 14063, 1, false, false),
    ('personal', 10000, 7, 14063, 1, false, false),
    ('personal', 10000, 7, 14063, 1, true, false),
    ('personal', 10000, 7, 13974, 1, true, false),
    ('personal', 12000, 7, 10193, 5, false, false),
    ('personal', 12000, 7, 10193, 5, true, false),
    ('personal', 14000, 29, 0, 1, false, false),
    ('personal', 14000, 29, 0, 1, true, false),
    ('personal', 15000, 7, 14063, 1, false, false),
    ('personal', 15000, 7, 14299, 1, true, false),
    ('personal', 18000, 7, 9702, 1, false, false),
    ('personal', 18000, 7, 9702, 1, true, false),
    ('personal', 20000, 7, 14063, 1, false, false),
    ('personal', 20000, 7, 14537, 1, true, false),
    ('personal', 22000, 26, 0, 4200, false, false),
    ('personal', 22000, 26, 0, 4200, true, false),
    ('personal', 25000, 7, 14063, 1, false, false),
    ('personal', 25000, 7, 14758, 1, true, false),
    ('personal', 26000, 7, 10194, 5, false, false),
    ('personal', 26000, 7, 10194, 5, true, false),
    ('personal', 30000, 7, 14063, 1, false, false),
    ('personal', 30000, 7, 14063, 1, false, false),
    ('personal', 30000, 7, 14063, 1, true, false),
    ('personal', 30000, 7, 14854, 1, true, false),
    ('personal', 34000, 29, 0, 2, false, false),
    ('personal', 34000, 29, 0, 2, true, false),
    ('personal', 40000, 7, 10195, 5, false, false),
    ('personal', 40000, 7, 10195, 5, true, false),
    ('personal', 46000, 26, 0, 4500, false, false),
    ('personal', 46000, 26, 0, 4500, true, false),
    ('personal', 50000, 7, 10196, 5, false, false),
    ('personal', 50000, 7, 10196, 5, true, false),
    ('personal', 54000, 29, 0, 3, false, false),
    ('personal', 54000, 29, 0, 3, true, false),
    ('personal', 60000, 7, 14063, 1, false, false),
    ('personal', 60000, 7, 14063, 1, true, false),
    ('personal', 63000, 26, 0, 4700, false, false),
    ('personal', 63000, 26, 0, 4700, true, false),
    ('personal', 70000, 7, 10197, 5, false, false),
    ('personal', 70000, 7, 10197, 5, true, false),
    ('personal', 72000, 7, 10198, 5, false, false),
    ('personal', 72000, 7, 10198, 5, true, false),
    ('personal', 74000, 29, 0, 4, false, false),
    ('personal', 74000, 29, 0, 4, true, false),
    ('personal', 78000, 26, 0, 5000, false, false),
    ('personal', 78000, 26, 0, 5000, true, false),
    ('personal', 82000, 7, 10199, 5, false, false),
    ('personal', 82000, 7, 10199, 5, true, false),
    ('personal', 84000, 29, 0, 5, false, false),
    ('personal', 84000, 29, 0, 5, true, false),
    ('personal', 86000, 26, 0, 5300, false, false),
    ('personal', 86000, 26, 0, 5300, true, false),
    ('personal', 90000, 7, 14063, 1, false, false),
    ('personal', 90000, 7, 14063, 1, true, false),
    ('personal', 92000, 7, 10730, 5, false, false),
    ('personal', 92000, 7, 10730, 5, true, false),
    ('personal', 94000, 29, 0, 6, false, false),
    ('personal', 94000, 29, 0, 6, true, false),
    ('personal', 98000, 7, 10731, 5, false, false),
    ('personal', 98000, 7, 10731, 5, true, false),
    ('personal', 102000, 26, 0, 5500, false, false),
    ('personal', 102000, 26, 0, 5500, true, false),
    ('personal', 104000, 29, 0, 7, false, false),
    ('personal', 104000, 29, 0, 7, true, false),
    ('personal', 106000, 7, 10732, 5, false, false),
    ('personal', 106000, 7, 10732, 5, true, false),
    ('personal', 110000, 7, 10189, 1, false, false),
    ('personal', 110000, 7, 10189, 1, true, false),
    ('personal', 114000, 29, 0, 8, false, false),
    ('personal', 114000, 29, 0, 8, true, false),
    ('personal', 118000, 26, 0, 5700, false, false),
    ('personal', 118000, 26, 0, 5700, true, false),
    ('personal', 124000, 29, 0, 9, false, false),
    ('personal', 124000, 29, 0, 9, true, false),
    ('personal', 126000, 7, 10188, 1, false, false),
    ('personal', 126000, 7, 10188, 1, true, false),
    ('personal', 134000, 29, 0, 10, false, false),
    ('personal', 134000, 29, 0, 10, true, false),
    ('personal', 146000, 26, 0, 5900, false, false),
    ('personal', 146000, 26, 0, 5900, true, false),
    ('personal', 150000, 7, 14063, 1, false, false),
    ('personal', 150000, 7, 14063, 1, true, false),
    ('personal', 160000, 26, 0, 6100, false, false),
    ('personal', 160000, 26, 0, 6100, true, false),
    ('personal', 174000, 26, 0, 6300, false, false),
    ('personal', 174000, 26, 0, 6300, true, false),
    ('personal', 180000, 7, 14063, 1, false, false),
    ('personal', 180000, 7, 14063, 1, true, false),
    ('personal', 186000, 26, 0, 6500, false, false),
    ('personal', 186000, 26, 0, 6500, true, false),
    ('personal', 200000, 7, 10187, 1, false, false),
    ('personal', 200000, 7, 10187, 1, true, false),
    ('personal', 214000, 26, 0, 6700, false, false),
    ('personal', 214000, 26, 0, 6700, true, false),
    ('personal', 226000, 7, 11440, 15, false, false),
    ('personal', 226000, 7, 11440, 15, true, false),
    ('personal', 240000, 26, 0, 7100, false, false),
    ('personal', 240000, 26, 0, 7100, true, false),
    ('personal', 260000, 26, 0, 1000, false, true),
    ('personal', 260000, 26, 0, 1000, true, true),
    ('personal', 280000, 26, 0, 1000, false, true),
    ('personal', 280000, 26, 0, 1000, true, true);

INSERT INTO public.diva_prizes
(type, points_req, item_type, item_id, quantity, gr, repeatable)
VALUES
    ('guild', 2, 7, 1026, 5, false, false),
    ('guild', 2, 7, 1026, 5, true, false),
    ('guild', 3, 7, 1026, 20, false, false),
    ('guild', 3, 7, 1026, 20, true, false),
    ('guild', 5, 7, 7456, 3, false, false),
    ('guild', 5, 7, 7456, 3, true, false),
    ('guild', 6, 7, 1026, 20, false, false),
    ('guild', 6, 7, 1026, 20, true, false),
    ('guild', 8, 7, 7457, 3, false, false),
    ('guild', 8, 7, 7457, 3, true, false),
    ('guild', 10, 7, 1026, 20, false, false),
    ('guild', 10, 7, 1026, 20, true, false),
    ('guild', 12, 7, 8940, 5, false, false),
    ('guild', 12, 7, 8941, 5, false, false),
    ('guild', 12, 7, 8943, 5, false, false),
    ('guild', 12, 7, 8946, 5, false, false),
    ('guild', 12, 7, 8940, 5, true, false),
    ('guild', 12, 7, 8941, 5, true, false),
    ('guild', 12, 7, 8943, 5, true, false),
    ('guild', 12, 7, 8946, 5, true, false),
    ('guild', 13, 26, 0, 1000, false, false),
    ('guild', 13, 26, 0, 1000, true, false),
    ('guild', 15, 7, 13692, 5, false, false),
    ('guild', 15, 7, 13693, 5, false, false),
    ('guild', 15, 7, 13692, 5, true, false),
    ('guild', 15, 7, 13693, 5, true, false),
    ('guild', 17, 26, 0, 2000, false, false),
    ('guild', 17, 26, 0, 2000, true, false),
    ('guild', 20, 28, 0, 1, false, false),
    ('guild', 20, 7, 7458, 5, false, false),
    ('guild', 20, 28, 0, 1, true, false),
    ('guild', 20, 7, 7458, 5, true, false),
    ('guild', 22, 7, 1026, 40, false, false),
    ('guild', 22, 7, 13692, 7, false, false),
    ('guild', 22, 7, 13693, 7, false, false),
    ('guild', 22, 7, 1026, 40, true, false),
    ('guild', 22, 7, 13692, 7, true, false),
    ('guild', 22, 7, 13693, 7, true, false),
    ('guild', 24, 7, 7463, 3, false, false),
    ('guild', 24, 7, 7463, 3, true, false),
    ('guild', 26, 26, 0, 3000, false, false),
    ('guild', 26, 26, 0, 3000, true, false),
    ('guild', 28, 7, 1026, 40, false, false),
    ('guild', 28, 7, 13692, 7, false, false),
    ('guild', 28, 7, 13693, 7, false, false),
    ('guild', 28, 7, 1026, 40, true, false),
    ('guild', 28, 7, 13692, 7, true, false),
    ('guild', 28, 7, 13693, 7, true, false),
    ('guild', 30, 7, 1026, 60, false, false),
    ('guild', 30, 7, 1026, 60, true, false),
    ('guild', 32, 7, 7462, 3, false, false),
    ('guild', 32, 7, 13692, 7, false, false),
    ('guild', 32, 7, 13693, 7, false, false),
    ('guild', 32, 7, 7462, 3, true, false),
    ('guild', 32, 7, 13692, 7, true, false),
    ('guild', 32, 7, 13693, 7, true, false),
    ('guild', 35, 7, 7464, 3, false, false),
    ('guild', 35, 7, 7464, 3, true, false),
    ('guild', 42, 7, 1026, 60, false, false),
    ('guild', 42, 7, 1026, 60, true, false),
    ('guild', 44, 7, 9710, 1, false, false),
    ('guild', 44, 7, 9710, 1, true, false),
    ('guild', 46, 7, 1026, 80, false, false),
    ('guild', 46, 7, 13692, 10, false, false),
    ('guild', 46, 7, 13693, 10, false, false),
    ('guild', 46, 7, 1026, 80, true, false),
    ('guild', 46, 7, 13692, 10, true, false),
    ('guild', 46, 7, 13693, 10, true, false),
    ('guild', 48, 7, 9709, 1, false, false),
    ('guild', 48, 7, 9709, 1, true, false),
    ('guild', 50, 7, 7456, 3, false, false),
    ('guild', 50, 7, 7456, 3, true, false),
    ('guild', 52, 7, 11387, 1, false, false),
    ('guild', 52, 7, 11387, 1, true, false),
    ('guild', 55, 7, 7457, 3, false, false),
    ('guild', 55, 7, 7457, 3, true, false),
    ('guild', 60, 7, 8945, 10, false, false),
    ('guild', 60, 7, 8945, 10, true, false),
    ('guild', 65, 7, 1026, 80, false, false),
    ('guild', 65, 7, 1026, 80, true, false),
    ('guild', 70, 7, 7458, 3, false, false),
    ('guild', 70, 7, 7458, 3, true, false),
    ('guild', 75, 7, 7463, 3, false, false),
    ('guild', 75, 7, 7463, 3, true, false),
    ('guild', 80, 7, 8945, 15, false, false),
    ('guild', 80, 7, 8945, 15, true, false),
    ('guild', 85, 7, 1026, 80, false, false),
    ('guild', 85, 7, 1026, 80, true, false),
    ('guild', 90, 7, 7462, 3, false, false),
    ('guild', 90, 7, 7462, 3, true, false),
    ('guild', 95, 7, 7464, 3, false, false),
    ('guild', 95, 7, 7464, 3, true, false),
    ('guild', 100, 26, 0, 50000, false, false),
    ('guild', 100, 26, 0, 50000, true, false);

END;
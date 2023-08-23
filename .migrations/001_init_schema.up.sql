BEGIN;

CREATE TABLE `ingredient` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `unit` varchar(50) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO ingredient VALUES
(1,"mourtarde à l'ancienne", "c.à.s"),
(2,"poivre", ""),
(3,"cheddar", "g"),
(4,"bière brune", "cl"),
(5,"pain de campagne non tranchée", ""),
(6,"jambon blanc", "tranche"),
(7,"oeufs", "");

CREATE TABLE `recipe` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(100) NOT NULL,
    `duration` int(11) NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO recipe VALUES (1, "Welsh traditionnel à la bière brune", 40); 

CREATE TABLE `recipe_description` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `recipe_id` int(11) NOT NULL,
    `description` TEXT NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO recipe_description VALUES
(1, 1, "Couper le cheddar en petits cubes. Couper le pain en tranches bien épaisses. Faire en sorte qu’au niveau largeur elles passent dans les plats à welsh. Dans la limite du possible garder la croute des tartines."),
(2, 1, "Verser quelques goutes de bière sur chaque tartine (vraiment quelques gouttes, il faut garder environ 20 cl pour le reste de la recette). Puis les badigeonner d’un peu de moutarde (environ 2 cuillères à soupe) et les placer dans le fond des plats."),
(3, 1, "Poser sur chaque tartine une tranche de jambon recourbée sur elle-même."),
(4, 1, "Dans une grande sauteuse, faire fondre les cubes de cheddar sans ajouter de matière grasse."),
(5, 1, "Mélanger constamment à l’aide d’une cuillère en bois (ça fait des fils, c’est normal tant que ça n’attache pas à la sauteuse)."),
(6, 1, "Quand la consistance est relativement homogène et qu’elle recouvre bien la cuillère quand on mélange, ajouter toute la bière et continuer à mélanger. Il faut remuer sans arrêt de façon à bien incorporer la bière au fromage."),
(7, 1, "Une fois le tout bien homogène et onctueux, ajouter le reste de la moutarde, un peu de poivre et remuer encore."),
(8, 1, "Quand la préparation est homogène, la verser dans les plats pour napper les tartines au jambon. Enfourner à 210°C (thermostat 7) pour une 10-12 minutes (il faut que ce soit bien doré)."),
(9, 1, "Pendant ce temps cuire les 4 oeufs au plat dans une poêle (il est aussi possible de les faire au four, directement sur les Welshs, mais la cuisson du jaune est très délicate. C'est pourquoi à la poêle c'est très bien)."),
(10, 1, "Quand le fromage est bien doré, sortir les plats du four, placer un oeuf au plat sur le dessus de chacun d’entre eux, tourner un coup de moulin à poivre et servir rapidement avec des frites et/ou de la salade verte. Et surtout une bière !");
 
CREATE TABLE `recipe_ingredient` (
    `recipe_id` int(11) NOT NULL,
    `ingredient_id` int(11) NOT NULL,
    `quantity` int(11) NOT NULL,
    UNIQUE (recipe_id, ingredient_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO recipe_ingredient VALUES
(1, 1, 4),
(1, 2, 1),
(1, 3, 800),
(1, 4, 25),
(1, 5, 25),
(1, 6, 1),
(1, 7, 4);

CREATE TABLE `favorite` (
    `user_id` int(11) NOT NULL,
    `recipe_id` int(11) NOT NULL,
    UNIQUE (user_id, recipe_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO favorite VALUES (1, 1); 

COMMIT;

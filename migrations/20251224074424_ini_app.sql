-- Create "schools" table
CREATE TABLE `schools` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  `name` text NOT NULL,
  `short_name` text NOT NULL,
  `address` text NULL
);
-- Create index "idx_schools_deleted_at" to table: "schools"
CREATE INDEX `idx_schools_deleted_at` ON `schools` (`deleted_at`);
-- Create "classrooms" table
CREATE TABLE `classrooms` (
  `id` integer NULL PRIMARY KEY AUTOINCREMENT,
  `created_at` datetime NULL,
  `updated_at` datetime NULL,
  `deleted_at` datetime NULL,
  `grade` text NOT NULL,
  `section` text NOT NULL,
  `school_id` integer NULL,
  CONSTRAINT `fk_classrooms_school` FOREIGN KEY (`school_id`) REFERENCES `schools` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- Create index "idx_classrooms_deleted_at" to table: "classrooms"
CREATE INDEX `idx_classrooms_deleted_at` ON `classrooms` (`deleted_at`);

-- name: GetAllChirps :many
SELECT *
FROM chirps
ORDER BY
  CASE 
    WHEN $1 = 'desc' THEN NULL
    ELSE created_at
  END ASC,
  CASE 
    WHEN $1 = 'desc' THEN created_at
    ELSE NULL
  END DESC;
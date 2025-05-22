// DOM Elements
const addProblemBtn = document.getElementById("addProblemBtn");
const problemForm = document.getElementById("problemForm");
const problemFormElement = document.getElementById("problemFormElement");
const problemsList = document.getElementById("problemsList");
const searchInput = document.getElementById("searchInput");
const tagFilter = document.getElementById("tagFilter");

// State
let problems = [];
let allTags = new Set();

// Event Listeners
addProblemBtn.addEventListener("click", () => {
  problemForm.style.display = "block";
});

problemFormElement.addEventListener("submit", handleSubmit);
searchInput.addEventListener("input", filterProblems);
tagFilter.addEventListener("change", filterProblems);

// Functions
function closeModal() {
  problemForm.style.display = "none";
  problemFormElement.reset();
}

async function handleSubmit(e) {
  e.preventDefault();

  const problem = {
    problem_id: parseInt(document.getElementById("problemId").value),
    title: document.getElementById("title").value,
    tags: document
      .getElementById("tags")
      .value.split(",")
      .map((tag) => tag.trim()),
  };

  try {
    const response = await fetch("/problems", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(problem),
    });

    if (!response.ok) throw new Error("Failed to create problem");

    closeModal();
    loadProblems();
  } catch (error) {
    console.error("Error:", error);
    alert("Failed to create problem");
  }
}

async function loadProblems() {
  try {
    const response = await fetch("/problems");
    if (!response.ok) throw new Error("Failed to fetch problems");

    problems = await response.json();
    updateTagsList();
    renderProblems();
  } catch (error) {
    console.error("Error:", error);
    alert("Failed to load problems");
  }
}

function updateTagsList() {
  allTags.clear();
  problems.forEach((problem) => {
    problem.tags.forEach((tag) => allTags.add(tag));
  });

  // Update tag filter options
  tagFilter.innerHTML = '<option value="">All Tags</option>';
  allTags.forEach((tag) => {
    const option = document.createElement("option");
    option.value = tag;
    option.textContent = tag;
    tagFilter.appendChild(option);
  });
}

function renderProblems(filteredProblems = problems) {
  problemsList.innerHTML = "";

  filteredProblems.forEach((problem) => {
    const card = document.createElement("div");
    card.className = "problem-card";

    card.innerHTML = `
            <h3>${problem.title}</h3>
            <p>Problem ID: ${problem.problem_id}</p>
            <div class="tags">
                ${problem.tags
                  .map((tag) => `<span class="tag">${tag}</span>`)
                  .join("")}
            </div>
            <div class="problem-actions">
                <button class="btn secondary" onclick="deleteProblem(${
                  problem.id
                })">Delete</button>
            </div>
        `;

    problemsList.appendChild(card);
  });
}

function filterProblems() {
  const searchTerm = searchInput.value.toLowerCase();
  const selectedTag = tagFilter.value;

  const filtered = problems.filter((problem) => {
    const matchesSearch =
      problem.title.toLowerCase().includes(searchTerm) ||
      problem.problem_id.toString().includes(searchTerm);
    const matchesTag = !selectedTag || problem.tags.includes(selectedTag);
    return matchesSearch && matchesTag;
  });

  renderProblems(filtered);
}

async function deleteProblem(id) {
  if (!confirm("Are you sure you want to delete this problem?")) return;

  try {
    const response = await fetch(`/problems/${id}`, {
      method: "DELETE",
    });

    if (!response.ok) throw new Error("Failed to delete problem");

    loadProblems();
  } catch (error) {
    console.error("Error:", error);
    alert("Failed to delete problem");
  }
}

// Initial load
loadProblems();

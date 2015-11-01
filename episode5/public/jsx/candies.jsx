var Root = React.createClass({
  getInitialState: function() {
    return {candyNames: []};
  },
  componentDidMount: function() {
    $.getJSON("/api/candies", function(data) {
      this.setState({candyNames: data["candies"]});
    }.bind(this));
  },
  render: function() {
    return (
      <div className="container candies">
        <div className="row">
          <CandiesList names={this.state.candyNames}/>
        </div>
        <div className="row">
          <h3>Create New Candy</h3>
          <CandyCreateForm />
        </div>
      </div>
    );
  }
});

var CandiesList = React.createClass({
  render: function() {
    var listElts = [];
    for(var i = 0; i < this.props.names.length; i++) {
      var name = this.props.names[i];
      console.log("name = " + name);
      listElts.push(<li key={name} className="list-group-item">{name}</li>);
    }
    return (
      <ul className="list-group">
        {listElts}
      </ul>
    );
  }
});

var CandyCreateForm = React.createClass({
  getInitialState: function() {
    return {candyName: ""};
  },
  handleChange: function(evt) {
    evt.preventDefault();
    this.setState({candyName: evt.target.value});
  },
  handleSubmit: function(e) {
    e.preventDefault();
    $.ajax({
      url: "/api/candy",
      data: JSON.stringify({"name": this.state.candyName}),
      method: "PUT",
      success: function() {
        this.setState({candyName: ""});
      }.bind(this)
    });
  },
  render: function() {
    return (
      <form onSubmit={this.handleSubmit}>
        <div className="form-group">
          <label for="candyName">Candy Name</label>
          <input type="text" className="form-control" value={this.state.candyName} id="candyName" placeholder="Candy Name" onChange={this.handleChange}/>
        </div>
        <button className="btn btn-default">Create</button>
      </form>
    );
  }
})

ReactDOM.render(<Root/>, document.getElementById("page-content"));
